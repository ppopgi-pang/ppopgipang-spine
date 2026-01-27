package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	userEntity "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type AuthService struct {
	jwtSecret        []byte
	jwtRefreshSecret []byte
}

func NewAuthService() *AuthService {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")

	if jwtSecret == "" || jwtRefreshSecret == "" {
		panic("JWT secrets이 필요합니다.")
	}

	return &AuthService{
		jwtSecret:        []byte(jwtSecret),
		jwtRefreshSecret: []byte(jwtRefreshSecret),
	}
}

func (a *AuthService) IssueAccessToken(user *userEntity.User) string {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"typ": "access",
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.jwtSecret)
	if err != nil {
		panic(err)
	}

	return signed
}

func (a *AuthService) IssueRefreshToken(user *userEntity.User) string {
	tokenID := uuid.NewString()

	claims := jwt.MapClaims{
		"sub": user.ID,
		"tid": tokenID,
		"typ": "refresh",
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.jwtRefreshSecret)
	if err != nil {
		panic(err)
	}

	// 🔒 서버에 refresh token 저장 (필수)
	user.RefreshToken = &tokenID
	// 예: a.userRepo.UpdateRefreshToken(user.ID, tokenID)

	return signed
}
