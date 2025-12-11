package server

import (
	"net/http"
)

type Protocol string

const (
	Http  Protocol = "http"
	Https Protocol = "https"
)

type ProxyRequest struct {
	Request  *http.Request
	Protocol Protocol
	Method   string
}

func (prh ProxyRequest) newProxyRequest(r *http.Request, p string) {
	prh.Request = r
	prh.Request.URL.Host = r.Host
	prh.Request.URL.Scheme = p
}
