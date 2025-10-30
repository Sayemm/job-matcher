package http

import (
	"net/http"

	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers"
)

type Server struct {
	jobHandler *handlers.Handler
	port       string
}

func NewServer(jobHandler *handlers.Handler, port string) *Server {
	return &Server{
		jobHandler: jobHandler,
		port:       port,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	s.jobHandler.RegisterRoutes(mux)

	addr := ":" + s.port
	http.ListenAndServe(addr, mux)
}
