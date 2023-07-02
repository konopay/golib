package httpclients

import (
	"crypto/tls"
	"crypto/x509"
	logs "github.com/cloudwego/kitex/pkg/klog"
	"github.com/konopay/golib/errs"
	"net/http"
	"net/url"
	"time"
)

type HttpClientOpts struct {
	Timeout                                 time.Duration
	ProxyUrl, RootCA, ClientCrt, PrivateKey string
}

// NewHttpClient 如有需要new一个http client，尽量一个server 只初始化一份
func NewHttpClient(opts *HttpClientOpts) (result *http.Client, err errs.Error) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("panic: error creating http client")
			err = errs.SystemError.ResetMsg("panic: error creating http client")
		}
	}()
	httpClient := &http.Client{}
	if opts == nil {
		logs.Info("httpclient config are empty")
		httpClient.Transport = &http.Transport{}
		return httpClient, nil
	}
	// set timeout
	httpClient.Timeout = getTimeout(opts.Timeout)
	// set transport: proxyUrl, rootCA, clientCert, privateKey...
	transport, err := getTransport(opts)
	if err != nil {
		return nil, err
	}
	httpClient.Transport = transport
	return httpClient, nil
}

func getTimeout(timeout time.Duration) time.Duration {
	if timeout == 0 {
		logs.Info("httpclient timeout set as default 10s")
		return DefaultTimeOut * time.Second
	}
	return timeout
}

func getTransport(transportConfig *HttpClientOpts) (*http.Transport, errs.Error) {
	transport := &http.Transport{}
	if transportConfig == nil {
		return transport, nil
	}
	// set proxy url
	if transportConfig.ProxyUrl != "" {
		parsedUrl, pErr := getProxyUrl(transportConfig.ProxyUrl)
		if pErr != nil {
			return nil, pErr
		}
		transport.Proxy = http.ProxyURL(parsedUrl)
		logs.Info("successfully added proxy")
	}

	rootCA := transportConfig.RootCA
	clientCrt := transportConfig.ClientCrt
	privateKey := transportConfig.PrivateKey

	tlsConf := &tls.Config{}
	if rootCA != "" {
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			logs.Error("loading system certs failed, err: %v", err)
		}
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		if ok := rootCAs.AppendCertsFromPEM([]byte(rootCA)); !ok {
			return nil, errs.SystemError.ResetMsg("adding gateway cert failed")
		}
		tlsConf.RootCAs = rootCAs
		logs.Info("successfully added root CA")
	}

	if clientCrt != "" && privateKey != "" {
		cert, loadErr := tls.X509KeyPair([]byte(clientCrt), []byte(privateKey))
		if loadErr != nil {
			logs.Error("loading ClientCrt and private failed, err: %v", loadErr)
			return nil, errs.SystemError.ResetMsg("adding client cert failed")
		}
		tlsConf.Certificates = []tls.Certificate{cert}
		logs.Info("successfully added client cert")
	}

	if tlsConf.RootCAs != nil || len(tlsConf.Certificates) > 0 {
		tlsConf.MinVersion = tls.VersionTLS12
		transport.TLSClientConfig = tlsConf
		logs.Info("successfully added tls config")
	}
	return transport, nil
}

func getProxyUrl(proxyUrl string) (*url.URL, errs.Error) {
	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			logs.Error("loading ProxyUrl failed, err: %v", err)
			return nil, errs.SystemError.ResetMsg("load proxy failed")
		}
		logs.Info("successfully added proxy")
		return proxy, nil
	}
	return nil, errs.SystemError.ResetMsg("invalid/empty proxyUrl")
}
