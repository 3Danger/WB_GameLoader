package server

import (
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run(port string, routing *http.ServeMux) error {
	s.srv = &http.Server{
		Addr:           ":" + port,
		Handler:        routing,
		TLSConfig:      nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.srv.ListenAndServe()
}
