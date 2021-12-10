package internalhttp

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type Server struct {
	logger Logger
	server *http.Server
}

type Logger interface {
	Debug(msg string)
	Warn(msg string)
	Info(msg string)
	Error(msg string)
}

func NewServer(logger Logger, address string, handler Handler) *Server {
	return &Server{
		logger,
		&http.Server{
			Addr:         address,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      handler.Router(),
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(err.Error())
		}
	}()

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Error("Server Shutdown.")

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Server Shutdown error:" + err.Error())
	}

	return nil
}
