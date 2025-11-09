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

// CORS Middleware - allows requests from frontend
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin (in production, specify exact origins)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow these HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allow these headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	s.jobHandler.RegisterRoutes(mux)
	s.resumeHandler.RegisterRoutes(mux)

	handler := corsMiddleware(mux)

	addr := ":" + s.port
	http.ListenAndServe(addr, handler)
}
