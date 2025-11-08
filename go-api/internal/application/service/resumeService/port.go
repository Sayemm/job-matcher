package resumeService

import (
	"github.com/Sayemm/job-matcher/go-api/internal/domain/entity"
	resumehandler "github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers/resumeHandler"
)

type Service interface {
	resumehandler.Service
}

type JobRepository interface {
	GetJobsByCluster(clusterID int, limit, offset int) ([]*entity.Job, error)
	CountByCluster(clusterID int) (int, error)
}
