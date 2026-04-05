package dto

import (
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type UserMeResponse struct {
	ID           int64     `json:"id" example:"1"`
	Email        string    `json:"email" example:"user@ppopgipang.com"`
	KakaoID      *string   `json:"kakaoId,omitempty" example:"1234567890"`
	Nickname     string    `json:"nickname" example:"뽑기팡유저"`
	ProfileImage *string   `json:"profileImage,omitempty" example:"https://cdn.ppopgi.me/profile.png"`
	IsAdmin      bool      `json:"isAdmin" example:"false"`
	MannerTemp   float64   `json:"mannerTemp" example:"36.5"`
	CreatedAt    time.Time `json:"createdAt" example:"2026-04-05T12:00:00Z"`
	UpdatedAt    time.Time `json:"updatedAt" example:"2026-04-05T12:00:00Z"`
}

func NewUserMeResponse(user entities.User) UserMeResponse {
	return UserMeResponse{
		ID:           user.ID,
		Email:        user.Email,
		KakaoID:      user.KakaoID,
		Nickname:     user.Nickname,
		ProfileImage: user.ProfileImage,
		IsAdmin:      user.IsAdmin,
		MannerTemp:   user.MannerTemp,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
