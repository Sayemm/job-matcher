package entity

import "time"

type Resume struct {
	ID         int       `json:"id"`
	Filename   string    `json:"filename"`
	Content    string    `json:"content"`
	ClusterID  *int      `json:"cluster_id"`
	UploadedAt time.Time `json:"uploaded_at"`
}
