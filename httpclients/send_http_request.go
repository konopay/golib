package httpclients

import (
	"compress/flate"
	"compress/gzip"
	"context"
	"github.com/andybalholm/brotli"
	logs "github.com/cloudwego/kitex/pkg/klog"
	"github.com/konopay/golib/errs"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type HttpRequest struct {
	Timeout time.Duration
	Method  string
	Url     string
	Headers map[string]string
	Body    string
	Proxy   string

	httpClient *http.Client // allow service to init their own HTTP Client, by call WithHttpClient
	opts       *HttpRequestOpts
}

type HttpRequestOpts struct {
	// NOTE: since the http header set/add/get will change the first letter with upper case.
	// That's not always meet our channel cases.
	// Please be aware: SendWithOrigHeaders mostly used for request channel.
	// if you want to send Http internally, you are still strongly suggested to use the Send one.
	// Since when you set key as `x` using this method,  you can't get `x` out of the header. get will only recognized `X`
	UseOriginalHeader bool
	// controls whether response bodies are logged in normal responses.
	// this prevents resource hogging when http resp data is huge
	PrintLog       bool // log all
	PrintRequest   bool // only log request body if true when PrintLog is false
	PrintReqHeader bool // only log request header if true when PrintLog is false
	PrintResponse  bool // only log response body+header if true when PrintLog is false

	// AutoDecompress option is set in upstream to determine if they want toolbox to perform auto decompression to the response
	AutoDecompress bool
}

var defaultOptions = &HttpRequestOpts{
	UseOriginalHeader: false,
	PrintLog:          true,
	PrintRequest:      false,
	PrintReqHeader:    false,
	PrintResponse:     false,
	AutoDecompress:    false,
}

func (o *HttpRequestOpts) Clone() *HttpRequestOpts {
	return &HttpRequestOpts{
		UseOriginalHeader: o.UseOriginalHeader,
		PrintLog:          o.PrintLog,
		PrintRequest:      o.PrintRequest,
		PrintReqHeader:    o.PrintReqHeader,
		PrintResponse:     o.PrintResponse,
		AutoDecompress:    o.AutoDecompress,
	}
}

func (r *HttpRequest) WithHttpClient(h *http.Client) {
	r.httpClient = h
}

func (r *HttpRequest) WithHttpRequestOps(h *HttpRequestOpts) {
	r.opts = h
}

type HttpResponse struct {
	HttpCode int
	Body     string
	Header   http.Header
}

func (r *HttpRequest) SendRequest(ctx context.Context) (*HttpResponse, errs.Error) {
	if r.opts == nil {
		r.opts = defaultOptions.Clone()
	}
	if r.Timeout == 0 {
		logs.CtxInfof(ctx, "http time set as default 30s")
		r.Timeout = DefaultTimeOut * time.Second
	}
	if r.httpClient == nil {
		r.httpClient = getHttpClientFromPool(r.Timeout)
	}
	rawHttpReq, err := http.NewRequestWithContext(ctx, r.Method, r.Url, strings.NewReader(r.Body))
	if err != nil || rawHttpReq == nil {
		logs.CtxErrorf(ctx, "make req error,err=%v", err)
		return nil, errs.SystemError
	}
	r.setHeaders(rawHttpReq)
	resp, rErr := r.httpClient.Do(rawHttpReq)
	if rErr != nil {
		respCode := 0
		if resp != nil {
			respCode = resp.StatusCode
		}
		if herr, ok := rErr.(net.Error); ok {
			if herr.Timeout() {
				logs.CtxErrorf(ctx, "HttpRequest timed out. response code %v", respCode)
				// the connection details is already logged during httptrace
				return &HttpResponse{HttpCode: respCode}, errs.HttpRemoteError
			}
		}
		logs.CtxErrorf(ctx, "HttpRequest failed. send request failed. error: %v", rErr)
		return &HttpResponse{HttpCode: respCode}, errs.New(errs.HttpRemoteError.Code(), rErr.Error())
	}
	httpBody := resp.Body
	httpCode := resp.StatusCode

	// if AutoDecompress is true in options, will proceed with decompressing
	if r.opts.AutoDecompress {
		logs.CtxInfof(ctx, "AutoDecompress option enabled, decompressing... Content-Encoding: %v", resp.Header.Get("Content-Encoding"))
		httpBody, err = decompressBody(httpBody, resp.Header.Get("Content-Encoding"))
		// if in gzip decompressing it returns nil, return the original response body
		if err != nil {
			logs.CtxErrorf(ctx, "http resp auto decompress gzip failed, code %v, body %v, err: %v", httpCode, resp.Body, err)
			return nil, errs.HttpRemoteError
		}
	} else {
		// support gzip read
		if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
			httpBody, err = gzip.NewReader(resp.Body)
			if err != nil {
				logs.CtxErrorf(ctx, "http resp unzip is failed, code %v, body %v, err: %v", httpCode, httpBody, err)
				return nil, errs.HttpRemoteError
			}
		}
	}
	httpResp, err := io.ReadAll(httpBody)
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logs.CtxErrorf(ctx, "Close body failed. error: %v.", err)
		}
	}()
	if err != nil {
		logs.CtxErrorf(ctx, "SendHttpRequest read response code %v, body %v, err: %v", httpCode, string(httpResp), err)
		return &HttpResponse{
			HttpCode: httpCode,
			Body:     string(httpResp),
		}, errs.HttpRemoteError
	}
	result := &HttpResponse{
		HttpCode: resp.StatusCode,
		Body:     string(httpResp),
		Header:   resp.Header,
	}
	return result, nil
}

func (r *HttpRequest) setHeaders(rawhttpReq *http.Request) {
	//add context content to http header
	var httpHeader = http.Header{}
	if r.opts.UseOriginalHeader {
		headerMap := make(map[string][]string)
		for key, value := range httpHeader { //put httpHeader content to request
			headerMap[key] = value
		}
		for key, value := range r.Headers {
			headerMap[key] = []string{value}
		}
		rawhttpReq.Header = headerMap
	} else {
		for key, value := range httpHeader {
			if len(value) > 0 {
				rawhttpReq.Header.Add(key, value[0])
			}
		}
		for key, value := range r.Headers {
			rawhttpReq.Header.Add(key, value)
		}
	}
}

// decompress body will decompress repsonse body according to Content-Encoding values. If there are multiple values
// decompressBody will decompress them in reverse order to get the original body. The only time it will return an error
// is when gzip unzip fails.
func decompressBody(body io.ReadCloser, encoding string) (io.ReadCloser, error) {
	// if more than one Content-Encoding value, we will handle with multiDecompress
	if strings.Contains(encoding, ", ") {
		return multiDecompress(body, encoding)
	}

	switch encoding {
	case "gzip":
		resp, err := gzip.NewReader(body)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "br":
		return ioutil.NopCloser(brotli.NewReader(body)), nil
	case "deflate":
		return flate.NewReader(body), nil
	}
	return body, nil
}

// decompresses the body in reverse order to get original value
func multiDecompress(body io.ReadCloser, encoding string) (io.ReadCloser, error) {
	encodings := strings.Split(encoding, ", ")
	resp := body
	for i := len(encodings) - 1; i >= 0; i-- {
		closer, err := decompressBody(resp, encodings[i])
		// if during multi decompression gzip returns an error, we return the original body
		if err != nil {
			return nil, err
		}
		resp = closer
	}
	return resp, nil
}
