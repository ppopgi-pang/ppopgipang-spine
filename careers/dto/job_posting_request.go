package dto

type JobPostingRequest struct {
	Title        string `json:"title" example:"백엔드 엔지니어"`
	Description  string `json:"description" example:"확장 가능한 API를 구축하고 운영합니다."`
	Department   string `json:"Department" example:"개발팀"`
	PositionType string `json:"position_type" example:"정규직"`
	Location     string `json:"location" example:"서울"`
}
