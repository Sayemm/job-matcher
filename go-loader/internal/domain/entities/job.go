package entities

import "time"

type Job struct {
	ID    int    `json:"id"`
	JobID string `json:"job_id"`

	CompanyName string `json:"company_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`

	RemoteAllowed   bool     `json:"remote_allowed"`
	ExperienceLevel string   `json:"experience_level"`
	MinSalary       *float64 `json:"min_salary,omitempty"`
	MaxSalary       *float64 `json:"max_salary,omitempty"`

	ClusterID *int `json:"cluster_id,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
