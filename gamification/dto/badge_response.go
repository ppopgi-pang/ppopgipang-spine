package dto

type NewBadgeResponse struct {
	ID             int64   `json:"id" example:"4001"`
	Code           *string `json:"code" example:"FIRST_CERTIFICATION"`
	Name           *string `json:"name" example:"첫 인증"`
	Description    *string `json:"description" example:"첫 인증을 완료했어요"`
	BadgeImageName *string `json:"badge_image_name" example:"badge_first_certification.png"`
}
