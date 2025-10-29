package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
)

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[ PROXY SERVER ] Request received at %s at %v\n", r.URL, time.Now().UTC())

		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host

		path := r.URL.Path
		r.URL.Path = strings.TrimPrefix(path, endpoint)

		proxyResp := fmt.Sprintf("[ PROXY SERVER ] Proxying request to %s at %s\n", r.URL, time.Now().UTC())

		// lit := "[ PROXY SERVER ]"
		// colorResp := color.Red(lit) + proxyResp
		// fmt.Println(colorResp)
		color.Red(proxyResp)

		proxy.ServeHTTP(w, r)
	}
}
