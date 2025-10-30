package entity

import "time"

type Job struct {
	ID              int       `db:"id" json:"id"`
	JobID           string    `db:"job_id" json:"job_id"`
	CompanyName     string    `db:"company_name" json:"company_name"`
	Title           string    `db:"title" json:"title"`
	Description     string    `db:"description" json:"description"`
	Location        string    `db:"location" json:"location"`
	RemoteAllowed   bool      `db:"remote_allowed" json:"remote_allowed"`
	ExperienceLevel string    `db:"experience_level" json:"experience_level"`
	MinSalary       *float64  `db:"min_salary" json:"min_salary,omitempty"`
	MaxSalary       *float64  `db:"max_salary" json:"max_salary,omitempty"`
	ClusterID       *int      `db:"cluster_id" json:"cluster_id,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
