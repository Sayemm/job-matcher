package handlers

import (
	"github.com/Sayemm/job-matcher/go-api/internal/application/dto"
	"github.com/Sayemm/job-matcher/go-api/internal/domain/entity"
)

type Service interface {
	GetJobByID(id int) (*entity.Job, error)
	GetJobs(page, pageSize int) (*dto.PaginatedResponse, error)
	GetJobsByCluster(clusterID int, page, pageSize int) (*dto.PaginatedResponse, error)
}
