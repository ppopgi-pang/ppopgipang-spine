package dto

type CommonResponse struct {
	ID      int64  `json:"id" example:"1"`
	Message string `json:"message" example:"요청이 정상적으로 처리되었습니다."`
}
