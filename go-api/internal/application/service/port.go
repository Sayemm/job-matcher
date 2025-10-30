package service

import (
	"github.com/Sayemm/job-matcher/go-api/internal/domain/entity"
	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers"
)

type Service interface {
	handlers.Service
}

type JobRepository interface {
	GetJobByID(id int) (*entity.Job, error)
	GetJobs(limit, offset int) ([]*entity.Job, error)
	GetJobsByCluster(clusterID int, limit, offset int) ([]*entity.Job, error)
	Count() (int, error)
	CountByCluster(clusterID int) (int, error)
}
