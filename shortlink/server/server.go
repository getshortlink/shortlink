// Package server provides the HTTP server for the ShortLink backend API.
package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/getshortlink/shortlink/server/link"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// TODO: Restrict allowed origins.
		AllowedOrigins:   []string{"https://*", "http://*", "chrome-extension://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s := &Server{
		router: r,
		server: &http.Server{
			Handler:      r,
			Addr:         fmt.Sprintf(":%v", DefaultPort),
			ReadTimeout:  DefaultReadTimeout,
			WriteTimeout: DefaultWriteTimeout,
		},
	}

	s.setHandlers()
	return s, nil
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownGracePeriod)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return
	}
}

func (s *Server) setHandlers() {
	s.router.Get("/health", s.handleHealth())
	s.router.Mount("/link", link.NewHandler().Router)
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
