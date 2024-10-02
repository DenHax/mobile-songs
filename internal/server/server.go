package server

import (
	"context"
	"net/http"

	"github.com/DenHax/mobile-songs/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(scfg config.ServerConfig, handler http.Handler) *Server {
	srv := new(Server)
	srv.httpServer = &http.Server{
		Addr:           scfg.Address,
		Handler:        handler,
		ReadTimeout:    scfg.ReadTimeout,
		WriteTimeout:   scfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	return srv
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
