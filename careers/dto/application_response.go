package dto

import "time"

type ApplicationResponse struct {
	ID                 int64                 `json:"id" example:"501"`
	JobPostingResponse ApplicationJobPosting `json:"job_posting_response"`
	Name               string                `json:"name" example:"김민수"`
	Email              string                `json:"email" example:"minsoo.kim@example.com"`
	Phone              *string               `json:"phone" example:"010-1234-5678"`
	ResumeFileName     *string               `json:"resume_file_name" example:"resume.pdf"`
	Memo               *string               `json:"memo" example:"백엔드 경력 5년"`
	Status             string                `json:"status" example:"접수"`
	CreatedAt          time.Time             `json:"created_at" example:"2024-03-01T10:20:30Z"`
}

type ApplicationJobPosting struct {
	ID    int64  `json:"id" example:"101"`
	Title string `json:"title" example:"백엔드 엔지니어"`
}

type ApplicationListResponse struct {
	Items []ApplicationResponse `json:"items"`
	Total int64                 `json:"total" example:"58"`
	Page  int                   `json:"page" example:"1"`
	Size  int                   `json:"size" example:"20"`
}
