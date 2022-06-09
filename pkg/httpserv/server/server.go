package server

import (
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func NewServer(port string, routing *http.ServeMux) *Server {
	serv := new(Server)
	serv.srv = &http.Server{
		Addr:           ":" + port,
		Handler:        routing,
		TLSConfig:      nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return serv
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
