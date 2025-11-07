package http

import (
	"net/http"

	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers/jobHandler"
	resumehandler "github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers/resumeHandler"
)

type Server struct {
	jobHandler    *jobHandler.Handler
	resumeHandler *resumehandler.Handler
	port          string
}

func NewServer(
	jobHandler *jobHandler.Handler,
	resumeHandler *resumehandler.Handler,
	port string,
) *Server {
	return &Server{
		jobHandler:    jobHandler,
		resumeHandler: resumeHandler,
		port:          port,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	s.jobHandler.RegisterRoutes(mux)
	s.resumeHandler.RegisterRoutes(mux)

	addr := ":" + s.port
	http.ListenAndServe(addr, mux)
}
