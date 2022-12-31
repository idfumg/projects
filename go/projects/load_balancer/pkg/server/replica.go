package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Replica struct {
	Url   *url.URL
	Proxy *httputil.ReverseProxy
}

func (s *Replica) Forward(w http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(w, r)
}
