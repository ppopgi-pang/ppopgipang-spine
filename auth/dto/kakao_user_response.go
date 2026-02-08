package dto

type KakaoUserResponse struct {
	ID           int64        `json:"id" example:"1234567890"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
}

type KakaoAccount struct {
	Email   string       `json:"email" example:"user@example.com"`
	Profile KakaoProfile `json:"profile"`
}

type KakaoProfile struct {
	Nickname string `json:"nickname" example:"뽑기팡유저"`
}
