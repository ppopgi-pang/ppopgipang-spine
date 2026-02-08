package dto

type StoreSearchResponse struct {
	Success bool            `json:"success" example:"true"`
	Data    []StoreResponse `json:"data"`
	Meta    Meta            `json:"meta"`
}
