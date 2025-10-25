package domain

import "github.com/Sayemm/job-matcher/go-loader/internal/domain/entities"

type JobRepository interface {
	SaveBath(jobs []*entities.Job) error
	Count() (int, error)
}
