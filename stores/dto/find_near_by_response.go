package dto

import "time"

type StoreFindNearByResponse struct {
	ID            int64             `json:"id"`
	Name          string            `json:"name"`
	Address       *string           `json:"address"`
	Region1       *string           `json:"region1"`
	Region2       *string           `json:"region2"`
	Latitude      float64           `json:"latitude"`
	Longitude     float64           `json:"longitude"`
	Phone         *string           `json:"phone"`
	AverageRating float32           `json:"average_rating"`
	Distance      int               `json:"distance"`
	Type          StoreTypeResponse `json:"type"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	RecentReview  *string           `json:"recent_review"`
	ReviewCount   int               `json:"review_count"`
}

type FindNearByDto struct {
	Success bool                      `json:"success"`
	Data    []StoreFindNearByResponse `json:"data"`
	Meta    Meta                      `json:"meta"`
}
