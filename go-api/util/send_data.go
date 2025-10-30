package util

import (
	"encoding/json"
	"net/http"

	"github.com/Sayemm/job-matcher/go-api/internal/application/dto"
)

func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, status int, message string, err error) {
	response := dto.ErrorResponse{
		Error:   err.Error(),
		Message: message,
	}
	SendJSON(w, status, response)
}
