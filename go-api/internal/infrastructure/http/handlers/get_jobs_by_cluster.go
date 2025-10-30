package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sayemm/job-matcher/go-api/util"
)

func (h *Handler) GetJobsByCluster(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter
	idStr := r.PathValue("id")
	fmt.Println(idStr)
	clusterID, err := strconv.Atoi(idStr)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid cluster ID", err)
		return
	}

	// Parse query parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	// Call service
	response, err := h.service.GetJobsByCluster(clusterID, page, pageSize)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to fetch jobs", err)
		return
	}

	util.SendJSON(w, http.StatusOK, response)
}
