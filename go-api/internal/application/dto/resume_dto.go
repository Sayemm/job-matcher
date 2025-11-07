package dto

type ResumeUploadResponse struct {
	Message  string `json:"message"`
	ResumeID int    `json:"resume_id"`
	Filename string `json:"filename"`
}

type ResumeMatchRequest struct {
	ResumeText string `json:"resume_text"`
}

type ResumeMatchResponse struct {
	ClusterID int     `json:"cluster_id"`
	Score     float64 `json:"score"`
}

type JobRecommendationResponse struct {
	Message        string      `json:"message"`
	ClusterID      int         `json:"cluster_id"`
	MatchScore     float64     `json:"match_score"`
	Jobs           interface{} `json:"jobs"`
	Page           int         `json:"page"`
	PageSize       int         `json:"page_size"`
	TotalInCluster int         `json:"total_in_cluster"`
}
