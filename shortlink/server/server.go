// Package server provides the HTTP server for the ShortLink backend API.
package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hellofresh/health-go/v5"
)

const (
	// Generally containers have a grace period of ten seconds to shutdown.
	ShutdownGracePeriod = 10 * time.Second

	// DefaultPort is the default port the server will listen on.
	DefaultPort = 8080

	// DefaultReadTimeout is the default timeout for reading the entire
	// request, including the body.
	DefaultReadTimeout = 5 * time.Second

	// DefaultWriteTimeout is the default timeout for writing the entire
	// response, including the body.
	DefaultWriteTimeout = 10 * time.Second
)

// Server is the HTTP server for the application.
type Server struct {
	router chi.Router
	server *http.Server
}

func NewServer() (*Server, error) {
	r := chi.NewRouter()

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Recover from panics without crashing the server, and log the panic
	// message.
	r.Use(middleware.Recoverer)

	s := &Server{
		router: r,
		server: &http.Server{
			Handler:      r,
			Addr:         fmt.Sprintf(":%v", DefaultPort),
			ReadTimeout:  DefaultReadTimeout,
			WriteTimeout: DefaultWriteTimeout,
		},
	}
	s.RegisterRoutes()

	return s, nil
}

// Start starts the server.
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Stop stops the server.
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownGracePeriod)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return
	}
}

// RegisterRoutes registers the routes for the server.
func (s *Server) RegisterRoutes() {
	s.router.Get("/health", s.handleHealth())
}

func (s *Server) handleHealth() http.HandlerFunc {
	handler, _ := health.New(
		health.WithComponent(health.Component{
			Name:    "shortlink",
			Version: "0.0.1",
		}),
	)

	return handler.Handler().ServeHTTP
}
