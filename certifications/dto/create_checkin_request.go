package dto

type CreateCheckInRequest struct {
	StoreId   int64    `json:"store_id" example:"2001"`
	Latitude  *float64 `json:"latitude" example:"37.4980"`
	Longitude *float64 `json:"longitude" example:"127.0276"`
	Rating    *string  `json:"rating" example:"good"`
	ReasonIds []int64  `json:"reason_ids" example:"1,2"`
}
