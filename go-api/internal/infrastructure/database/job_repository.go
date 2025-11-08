package database

import (
	"github.com/Sayemm/job-matcher/go-api/internal/application/service/jobService"
	"github.com/Sayemm/job-matcher/go-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type JobRepository interface {
	jobService.JobRepository
}

type jobRepo struct {
	db *sqlx.DB
}

func NewJobRepo(db *sqlx.DB) JobRepository {
	return &jobRepo{
		db: db,
	}
}

func (r *jobRepo) GetJobByID(id int) (*entity.Job, error) {
	var job entity.Job
	query := "SELECT * FROM jobs WHERE id = $1"
	err := r.db.Get(&job, query, id)
	return &job, err
}

func (r *jobRepo) GetJobs(limit, offset int) ([]*entity.Job, error) {
	var jobs []*entity.Job
	query := "SELECT * FROM jobs ORDER BY id LIMIT $1 OFFSET $2"
	err := r.db.Select(&jobs, query, limit, offset)
	return jobs, err
}

func (r *jobRepo) GetJobsByCluster(clusterID int, limit, offset int) ([]*entity.Job, error) {
	var jobs []*entity.Job
	query := "SELECT * FROM jobs ORDER BY id LIMIT $1 OFFSET $2"
	err := r.db.Select(&jobs, query, limit, offset)
	return jobs, err
}

func (r *jobRepo) Count() (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM jobs")
	return count, err
}

func (r *jobRepo) CountByCluster(clusterID int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM jobs WHERE cluster_id = $1"
	err := r.db.Get(&count, query, clusterID)
	return count, err
}
