package dto

type StoreNearestResponse struct {
	ID            int64   `json:"id" example:"2000"`
	Name          string  `json:"name" example:"뽑기팡 매장"`
	Latitude      float64 `json:"latitude" example:"37.5000"`
	Longitude     float64 `json:"longitude" example:"127.0350"`
	AverageRating float32 `json:"average_rating" example:"4.0"`
	Distance      int     `json:"distance" example:"120"`
	ThumbnailName *string `json:"thumbnail_name" example:"store_thumbnail_2000.jpg"`
	ReviewCount   int     `json:"review_count" example:"100"`
}

type FindNearestStoreResponse struct {
	Success bool                 `json:"success" example:"true"`
	Data    StoreNearestResponse `json:"data"`
}
