package dto

import "time"

type ApplicationResponse struct {
	ID                 int64                 `json:"id"`
	JobPostingResponse ApplicationJobPosting `json:"job_posting_response"`
	Name               string                `json:"name"`
	Email              string                `json:"email"`
	Phone              *string               `json:"phone"`
	ResumeFileName     *string               `json:"resume_file_name"`
	Memo               *string               `json:"memo"`
	Status             string                `json:"status"`
	CreatedAt          time.Time             `json:"created_at"`
}

type ApplicationJobPosting struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type ApplicationListResponse struct {
	Items []ApplicationResponse `json:"items"`
	Total int64                 `json:"total"`
	Page  int                   `json:"page"`
	Size  int                   `json:"size"`
}
