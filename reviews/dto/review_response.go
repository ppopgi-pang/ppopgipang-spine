package dto

import "time"

type ReviewResponse struct {
	ID        int64     `json:"id"`
	Rating    int       `json:"rating"`
	Content   *string   `json:"content"`
	Images    []string  `json:"images"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
