package dto

type AdminUserRequest struct {
	Email    string `json:"email" example:"admin@ppopgipang.com"`
	Nickname string `json:"nickname" example:"관리자"`
	Password string `json:"password" example:"P@ssw0rd123!"`
}
