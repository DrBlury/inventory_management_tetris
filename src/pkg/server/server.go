package server

import (
	"context"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *Config, mux http.Handler) *Server {
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: mux,
	}

	return &Server{
		server: server,
	}
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctxShutDown context.Context) error {
	return s.server.Shutdown(ctxShutDown)
}
