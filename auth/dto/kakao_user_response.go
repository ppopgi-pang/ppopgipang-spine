package dto

type KakaoUserResponse struct {
	ID           int64        `json:"id"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
}

type KakaoAccount struct {
	Email   string       `json:"email"`
	Profile KakaoProfile `json:"profile"`
}

type KakaoProfile struct {
	Nickname string `json:"nickname"`
}
