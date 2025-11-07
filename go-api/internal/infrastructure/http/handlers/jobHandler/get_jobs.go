package jobHandler

import (
	"net/http"
	"strconv"

	"github.com/Sayemm/job-matcher/go-api/util"
)

func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pagesize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	response, err := h.service.GetJobs(page, pagesize)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to fetch jobs", err)
		return
	}

	util.SendJSON(w, http.StatusOK, response)
}
