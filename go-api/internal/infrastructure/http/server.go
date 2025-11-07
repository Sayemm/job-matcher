package http

import (
	"net/http"

	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers/jobHandler"
)

type Server struct {
	jobHandler *jobHandler.Handler
	port       string
}

func NewServer(jobHandler *jobHandler.Handler, port string) *Server {
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
