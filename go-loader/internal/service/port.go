package service

import "github.com/Sayemm/job-matcher/go-loader/internal/domain"

// JobRepository defines what database operations we need
type JobRepository interface {
	SaveBatch(jobs []*domain.Job) error
	Count() (int, error)
}

// CSVReader defines what CSV operations we need
type CSVReader interface {
	ReadInBatches(batchSize int, callback func([]*domain.Job) error) error
}

type Service interface {
	LoadJobs() error
}
