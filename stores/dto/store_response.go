package dto

import "time"

type StoreTypeResponse struct {
	ID          int64   `json:"id" example:"1"`
	Name        string  `json:"name" example:"카페"`
	Description *string `json:"description" example:"커피와 디저트"`
}

type StoreResponse struct {
	ID            int64             `json:"id" example:"2000"`
	Name          string            `json:"name" example:"뽑기팡 매장"`
	Address       *string           `json:"address" example:"서울 강남구 테헤란로 10"`
	Region1       *string           `json:"region1" example:"서울"`
	Region2       *string           `json:"region2" example:"강남구"`
	Latitude      float64           `json:"latitude" example:"37.5000"`
	Longitude     float64           `json:"longitude" example:"127.0350"`
	Phone         *string           `json:"phone" example:"02-555-1234"`
	AverageRating float32           `json:"average_rating" example:"4.0"`
	Distance      int               `json:"distance" example:"120"`
	Type          StoreTypeResponse `json:"type"`
	ThumbnailName *string           `json:"thumbnail_name" example:"store_thumbnail_2000.jpg"`
	CreatedAt     time.Time         `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt     time.Time         `json:"updated_at" example:"2024-02-01T00:00:00Z"`
}

type StoreSummaryResponse struct {
	ID            int64    `json:"id" example:"2000"`
	Name          string   `json:"name" example:"뽑기팡 매장"`
	AverageRating float32  `json:"average_rating" example:"4.0"`
	ImageNames    []string `json:"image_names"`
	ReviewCount   int      `json:"review_count" example:"100"`
}

type StoreDetailResponse struct {
	IsBookmark                bool                       `json:"is_bookmark"`
	StoreOpeningHourResponses []StoreOpeningHourResponse `json:"store_opening_hour_responses"`
	Phone                     *string                    `json:"phone"`
	StoreFacilityResponse     StoreFacilityResponse      `json:"store_facility_response"`
}
