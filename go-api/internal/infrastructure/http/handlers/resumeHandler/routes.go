package resumehandler

import (
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /api/resume/upload-and-match", http.HandlerFunc(h.UploadAndMatch))
}
