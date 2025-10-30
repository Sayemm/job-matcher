package repository

import "github.com/Sayemm/job-matcher/go-api/internal/domain/entity"

type JobRepository interface {
	GetByID(id int) (*entity.Job, error)
	GetAll(limit, offset int) ([]*entity.Job, error)
	GetByCluster(clusterID int, limit, offset int) ([]*entity.Job, error)
	Count() (int, error)
	CountByCluster(clusterID int) (int, error)
}
