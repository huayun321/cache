package http

import (
	"github.com/huayun321/cache/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}

func NewServer(cache cache.Cache) *Server {
	return &Server{Cache: cache}
}
