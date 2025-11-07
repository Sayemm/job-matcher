package resumeservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Sayemm/job-matcher/go-api/internal/application/dto"
)

type service struct {
	jobRepository JobRepository
	mlServiceURL  string
}

func NewResumeService(jobRepository JobRepository, mlServiceURL string) Service {
	return &service{
		jobRepository: jobRepository,
		mlServiceURL:  mlServiceURL,
	}
}

func (s *service) ExtractTextFromFile(file multipart.File, filename string) (string, error) {
	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// For now, assume it's a text file
	// Later we will add PDF parsing
	text := string(content)

	return text, nil
}
func (s *service) FindMatchingCluster(resumeText string) (*dto.ResumeMatchResponse, error) {
	// Prepare request
	request := dto.ResumeMatchRequest{
		ResumeText: resumeText,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Call ML service
	url := s.mlServiceURL + "/match"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to call ML service: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var matchResponse dto.ResumeMatchResponse
	if err := json.NewDecoder(resp.Body).Decode(&matchResponse); err != nil {
		return nil, fmt.Errorf("failed to parse ML response: %w", err)
	}

	return &matchResponse, nil
}
func (s *service) GetRecommendedJobs(clusterID int, page, pageSize int) (*dto.JobRecommendationResponse, error) {
	// Validate
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	// Get jobs from cluster
	jobs, err := s.jobRepository.GetJobsByCluster(clusterID, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %w", err)
	}

	// Get total count
	totalCount, err := s.jobRepository.CountByCluster(clusterID)
	if err != nil {
		return nil, fmt.Errorf("failed to count jobs: %w", err)
	}

	// Build response
	response := &dto.JobRecommendationResponse{
		Message:        fmt.Sprintf("Found %d matching jobs in your cluster", totalCount),
		ClusterID:      clusterID,
		Jobs:           jobs,
		Page:           page,
		PageSize:       pageSize,
		TotalInCluster: totalCount,
	}

	return response, nil
}
