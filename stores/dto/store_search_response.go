package dto

type StoreSearchResponse struct {
	Success bool            `json:"success"`
	Data    []StoreResponse `json:"data"`
	Meta    Meta            `json:"meta"`
}
