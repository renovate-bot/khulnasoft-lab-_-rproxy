package proxy

import (
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(upstreamURL *url.URL) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.Out.URL.Host = upstreamURL.Host
			r.Out.URL.Scheme = upstreamURL.Scheme
			r.Out.Host = upstreamURL.Host
		},
	}
}
