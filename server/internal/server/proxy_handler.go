package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type ProxyRequestHandler struct {
	temp string
}

func (prh ProxyRequestHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[PROXY SERVER] request received at %s at %s \n forwading request...\n", r.URL, time.Now().UTC())

	if r.Host == "" {
		http.Error(w, "Host Not Found", http.StatusNotFound)
		return
	}

	r.URL.Host = r.Host
	r.URL.Scheme = "http"

	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	fmt.Printf("[PROXY SERVER] response received from %s \n", r.Host)

	//copy response to http writer
	defer resp.Body.Close() // defer till done copying response

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
