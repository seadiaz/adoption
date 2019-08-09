package adapters

import (
	"net/http"
)

// HTTPServer ...
type HTTPServer interface {
	ListenAndServe() error
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
	s.HTTPServer.ListenAndServe()
}
