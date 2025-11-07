package jobHandler

import (
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/jobs", http.HandlerFunc(h.GetJobs))
	mux.Handle("GET /api/jobs/{id}", http.HandlerFunc(h.GetJobById))
	mux.Handle("GET /api/jobs/cluster/{id}", http.HandlerFunc(h.GetJobsByCluster))
}
