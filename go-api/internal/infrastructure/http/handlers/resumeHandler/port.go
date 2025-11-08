package resumeHandler

import (
	"mime/multipart"

	"github.com/Sayemm/job-matcher/go-api/internal/application/dto"
)

type Service interface {
	ExtractTextFromFile(file multipart.File, filename string) (string, error)
	FindMatchingCluster(resumeText string) (*dto.ResumeMatchResponse, error)
	GetRecommendedJobs(clusterID int, page, pageSize int) (*dto.JobRecommendationResponse, error)
}
