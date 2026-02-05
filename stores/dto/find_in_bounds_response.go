package dto

import "time"

type StoreInBoundsResponse struct {
	ID            int64             `json:"id" example:"2002"`
	Name          string            `json:"name" example:"뽑기팡"`
	Address       *string           `json:"address" example:"서울 마포구 마포대로 45"`
	Region1       *string           `json:"region1" example:"서울"`
	Region2       *string           `json:"region2" example:"마포구"`
	Latitude      float64           `json:"latitude" example:"37.5410"`
	Longitude     float64           `json:"longitude" example:"126.9510"`
	Phone         *string           `json:"phone" example:"02-987-6543"`
	AverageRating float32           `json:"average_rating" example:"4.2"`
	Type          StoreTypeResponse `json:"type"`
	CreatedAt     time.Time         `json:"created_at" example:"2024-01-12T09:00:00Z"`
	UpdatedAt     time.Time         `json:"updated_at" example:"2024-02-07T14:30:00Z"`
	RecentReview  *string           `json:"recent_review" example:"뽑기가 잘 뽑혀요"`
	ReviewCount   int               `json:"review_count" example:"18"`
}

type FindInBoundsDto struct {
	Success bool                    `json:"success" example:"true"`
	Data    []StoreInBoundsResponse `json:"data"`
	Meta    Meta                    `json:"meta"`
}
