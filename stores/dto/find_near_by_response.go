package dto

import "time"

type StoreFindNearByResponse struct {
	ID            int64             `json:"id" example:"2001"`
	Name          string            `json:"name" example:"뽑기팡"`
	Address       *string           `json:"address" example:"서울 강남구 강남대로 123"`
	Region1       *string           `json:"region1" example:"서울"`
	Region2       *string           `json:"region2" example:"강남구"`
	Latitude      float64           `json:"latitude" example:"37.4980"`
	Longitude     float64           `json:"longitude" example:"127.0276"`
	Phone         *string           `json:"phone" example:"02-123-4567"`
	AverageRating float32           `json:"average_rating" example:"4.5"`
	Distance      int               `json:"distance" example:"350"`
	Type          StoreTypeResponse `json:"type"`
	CreatedAt     time.Time         `json:"created_at" example:"2024-01-10T08:15:00Z"`
	UpdatedAt     time.Time         `json:"updated_at" example:"2024-02-05T12:00:00Z"`
	RecentReview  *string           `json:"recent_review" example:"뽑기가 잘 뽑혀요"`
	ReviewCount   int               `json:"review_count" example:"27"`
}

type FindNearByDto struct {
	Success bool                      `json:"success" example:"true"`
	Data    []StoreFindNearByResponse `json:"data"`
	Meta    Meta                      `json:"meta"`
}
