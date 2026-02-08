package dto

type KakaoTokenResponse struct {
	AccessToken  string `json:"access_token" example:"kakao-access-token-sample"`
	TokenType    string `json:"token_type" example:"bearer"`
	RefreshToken string `json:"refresh_token" example:"kakao-refresh-token-sample"`
	ExpiresIn    int    `json:"expires_in" example:"21599"`
}
