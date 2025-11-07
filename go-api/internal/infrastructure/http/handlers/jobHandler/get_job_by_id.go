package jobHandler

import (
	"net/http"
	"strconv"

	"github.com/Sayemm/job-matcher/go-api/util"
)

func (h *Handler) GetJobById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid job ID", err)
		return
	}

	job, err := h.service.GetJobByID(id)
	if err != nil {
		util.SendError(w, http.StatusNotFound, "Job not found", err)
		return
	}

	util.SendJSON(w, http.StatusOK, job)
}
