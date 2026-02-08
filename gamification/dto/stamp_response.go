package dto

type StampResponse struct {
	ID        int64   `json:"id" example:"3001"`
	ImageName *string `json:"image_name" example:"stamp_3001.png"`
	StoreName *string `json:"store_name" example:"뽑기팡 강남점"`
}
