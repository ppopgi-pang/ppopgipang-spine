package dto

import "time"

type JobPostingResponse struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Department       string    `json:"department"`
	PositionType     string    `json:"position_type"`
	Location         string    `json:"location"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `jsonœ:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ApplicationCount int64     `json:"application_count"`
}

type JobPostingListResponse struct {
	Items []JobPostingResponse `json:"item"`
	Total int64                `json:"total"`
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
}
