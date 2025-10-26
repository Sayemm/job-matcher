package repository

import (
	"fmt"

	"github.com/Sayemm/job-matcher/go-loader/internal/domain"
	"github.com/Sayemm/job-matcher/go-loader/internal/service"
	"github.com/jmoiron/sqlx"
)

type JobRepository interface {
	service.JobRepository
}

type jobRepository struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) JobRepository {
	return &jobRepository{
		db: db,
	}
}

func (j *jobRepository) SaveBatch(jobs []*domain.Job) error {
	if len(jobs) == 0 {
		return nil
	}

	tx, err := j.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO jobs (
			job_id, company_name, title, description, location,
			remote_allowed, experience_level, min_salary, max_salary
		) VALUES (
			:job_id, :company_name, :title, :description, :location,
			:remote_allowed, :experience_level, :min_salary, :max_salary
		)
		ON CONFLICT (job_id) DO NOTHING
	`

	for _, job := range jobs {
		_, err := tx.NamedExec(query, job)
		if err != nil {
			return fmt.Errorf("failed to insert job: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
func (j *jobRepository) Count() (int, error) {
	var count int
	err := j.db.Get(&count, "SELECT COUNT(*) FROM jobs")
	return count, err
}
