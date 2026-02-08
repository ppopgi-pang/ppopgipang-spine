package dto

type JobPostingModifyRequest struct {
	Title        *string `json:"title" example:"시니어 백엔드 엔지니어"`
	Description  *string `json:"description" example:"대규모 트래픽 환경에서 API를 설계하고 운영합니다."`
	Department   *string `json:"Department" example:"플랫폼팀"`
	PositionType *string `json:"position_type" example:"정규직"`
	Location     *string `json:"location" example:"서울"`
	IsActive     *bool   `json:"is_active" example:"true"`
}
