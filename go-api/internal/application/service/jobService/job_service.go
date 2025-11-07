package jobservice

import (
	"fmt"

	"github.com/Sayemm/job-matcher/go-api/internal/application/dto"
	"github.com/Sayemm/job-matcher/go-api/internal/domain/entity"
)

type service struct {
	jobRepository JobRepository
}

func NewJobService(jobRepository JobRepository) Service {
	return &service{
		jobRepository: jobRepository,
	}
}

func (svc *service) GetJobByID(id int) (*entity.Job, error) {
	if id < 1 {
		return nil, fmt.Errorf("invalid job ID: %d", id)
	}
	job, err := svc.jobRepository.GetJobByID(id)
	if err != nil {
		return nil, fmt.Errorf("job not found: %w", err)
	}
	return job, nil
}
func (svc *service) GetJobs(page, pageSize int) (*dto.PaginatedResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	jobs, err := svc.jobRepository.GetJobs(pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %w", err)
	}

	totalCount, err := svc.jobRepository.Count()
	if err != nil {
		return nil, fmt.Errorf("failed to count jobs: %w", err)
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	response := &dto.PaginatedResponse{
		Data:       jobs,
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return response, nil
}
func (svc *service) GetJobsByCluster(clusterID int, page, pageSize int) (*dto.PaginatedResponse, error) {
	if clusterID < 0 {
		return nil, fmt.Errorf("invalid cluster ID: %d", clusterID)
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	jobs, err := svc.jobRepository.GetJobsByCluster(clusterID, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs by cluster: %w", err)
	}

	totalCount, err := svc.jobRepository.CountByCluster(clusterID)
	if err != nil {
		return nil, fmt.Errorf("failed to count jobs in cluster: %w", err)
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	response := &dto.PaginatedResponse{
		Data:       jobs,
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return response, nil
}
