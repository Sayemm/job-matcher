package resumeService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

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

// ExtractTextFromFile extracts text from uploaded file
func (s *service) ExtractTextFromFile(file multipart.File, filename string) (string, error) {
	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Check file type by extension
	filename = strings.ToLower(filename)

	if strings.HasSuffix(filename, ".pdf") {
		// Send PDF to Python for parsing
		return s.extractTextFromPDF(content)
	} else if strings.HasSuffix(filename, ".txt") {
		// Plain text file
		return string(content), nil
	} else {
		return "", fmt.Errorf("unsupported file type: %s (only .pdf and .txt supported)", filename)
	}
}

// extractTextFromPDF sends PDF to Python service for text extraction
func (s *service) extractTextFromPDF(pdfContent []byte) (string, error) {
	// Create request to ml-matcher's PDF parsing endpoint
	url := s.mlServiceURL + "/parse-pdf"

	resp, err := http.Post(url, "application/pdf", bytes.NewBuffer(pdfContent))
	if err != nil {
		return "", fmt.Errorf("failed to call PDF parser: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("PDF parsing failed with status %d", resp.StatusCode)
	}

	// Parse response
	var result struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse PDF response: %w", err)
	}

	return result.Text, nil
}

// FindMatchingCluster sends resume to ML service
func (s *service) FindMatchingCluster(resumeText string) (*dto.ResumeMatchResponse, error) {
	request := dto.ResumeMatchRequest{
		ResumeText: resumeText,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := s.mlServiceURL + "/match"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to call ML service: %w", err)
	}
	defer resp.Body.Close()

	var matchResponse dto.ResumeMatchResponse
	if err := json.NewDecoder(resp.Body).Decode(&matchResponse); err != nil {
		return nil, fmt.Errorf("failed to parse ML response: %w", err)
	}

	return &matchResponse, nil
}

// GetRecommendedJobs gets jobs from matched cluster
func (s *service) GetRecommendedJobs(clusterID int, page, pageSize int) (*dto.JobRecommendationResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	jobs, err := s.jobRepository.GetJobsByCluster(clusterID, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %w", err)
	}

	totalCount, err := s.jobRepository.CountByCluster(clusterID)
	if err != nil {
		return nil, fmt.Errorf("failed to count jobs: %w", err)
	}

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
