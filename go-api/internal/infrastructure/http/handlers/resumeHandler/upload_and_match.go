package resumeHandler

import (
	"net/http"

	"github.com/Sayemm/job-matcher/go-api/util"
)

func (h *Handler) UploadAndMatch(w http.ResponseWriter, r *http.Request) {
	// Step 1: Parse multipart form (max 10MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Failed to parse form", err)
		return
	}

	// Step 2: Get file from form
	file, header, err := r.FormFile("resume")
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Resume file is required", err)
		return
	}
	defer file.Close()

	// Step 3: Extract text from file
	resumeText, err := h.service.ExtractTextFromFile(file, header.Filename)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to extract text", err)
		return
	}

	// Step 4: Find matching cluster (call ML service)
	matchResult, err := h.service.FindMatchingCluster(resumeText)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to match resume", err)
		return
	}

	// Step 5: Get recommended jobs from that cluster
	recommendations, err := h.service.GetRecommendedJobs(matchResult.ClusterID, 1, 20)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to get recommendations", err)
		return
	}

	// Step 6: Add match score to response
	recommendations.MatchScore = matchResult.Score

	// Step 7: Send response
	util.SendJSON(w, http.StatusOK, recommendations)
}
