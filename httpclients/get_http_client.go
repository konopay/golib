package httpclients

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	httpClientPool          = make(map[time.Duration]*http.Client)
	httpClientPoolWithProxy = make(map[string]*http.Client)
	rwLock                  = sync.RWMutex{}
)

func getHttpClientFromPool(timeout time.Duration) *http.Client {
	if httpClientPool == nil {
		httpClientPool = make(map[time.Duration]*http.Client)
	}
	rwLock.RLock()
	httpClient, found := httpClientPool[timeout]
	rwLock.RUnlock()
	if !found {
		httpClient = &http.Client{
			Timeout: timeout,
		}
		rwLock.Lock()
		httpClientPool[timeout] = httpClient
		rwLock.Unlock()
	}
	return httpClient
}

func getHttpProxyClientFromPool(ctx context.Context, timeout time.Duration, proxy string) *http.Client {
	identity := fmt.Sprintf("%v_%v", timeout, proxy)
	if httpClientPoolWithProxy == nil {
		httpClientPoolWithProxy = make(map[string]*http.Client)
	}
	rwLock.RLock()
	httpClient, found := httpClientPoolWithProxy[identity]
	rwLock.RUnlock()
	if !found {
		httpClient = &http.Client{
			Timeout: timeout,
		}
		if proxy != "" {
			proxy, _ := url.Parse(proxy)
			httpClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		}
		rwLock.Lock()
		httpClientPoolWithProxy[identity] = httpClient
		rwLock.Unlock()
	}
	return httpClient
}
