package dto

import "time"

type JobPostingResponse struct {
	ID               int64     `json:"id" example:"101"`
	Title            string    `json:"title" example:"백엔드 엔지니어"`
	Description      string    `json:"description" example:"확장 가능한 API를 구축하고 운영합니다."`
	Department       string    `json:"department" example:"개발팀"`
	PositionType     string    `json:"position_type" example:"정규직"`
	Location         string    `json:"location" example:"서울"`
	IsActive         bool      `json:"is_active" example:"true"`
	CreatedAt        time.Time `json:"created_at" example:"2024-01-02T15:04:05Z"`
	UpdatedAt        time.Time `json:"updated_at" example:"2024-02-10T09:30:00Z"`
	ApplicationCount int64     `json:"application_count" example:"12"`
}

type JobPostingListResponse struct {
	Items []JobPostingResponse `json:"item"`
	Total int64                `json:"total" example:"120"`
	Page  int                  `json:"page" example:"1"`
	Size  int                  `json:"size" example:"20"`
}
