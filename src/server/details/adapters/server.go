package adapters

import (
	"context"
	"net/http"

	"github.com/golang/glog"
)

// HTTPServer ...
type HTTPServer interface {
	ListenAndServe() error
	Close() error
	Shutdown(ctx context.Context) error
}

// Server ...
type Server struct {
	HTTPServer HTTPServer
	Router     Router
}

// Router ...
type Router interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route
}

// Route ...
type Route interface {
	Methods(methods ...string) Route
}

// CreateServer ...
func CreateServer(server HTTPServer, router Router) Server {
	return Server{
		HTTPServer: server,
		Router:     router,
	}
}

// Run ...
func (s *Server) Run() {
	err := s.HTTPServer.ListenAndServe()
	glog.Error(err)
}

// Close ...
func (s *Server) Close() {
	s.HTTPServer.Close()
}

// Shutdown ...
func (s *Server) Shutdown() {
	s.HTTPServer.Shutdown(nil)
}
