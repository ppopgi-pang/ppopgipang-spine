package controller

import (
	"context"
	"errors"
	"time"

	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/service"
	userEntity "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (a *AuthController) KakaoCallback(ctx context.Context, query query.Values, spineCtx spine.Ctx) (httpx.Redirect, error) {
	userAny, ok := spineCtx.Get("auth.user")
	if !ok {
		panic(httperr.Unauthorized("인증 정보가 없습니다."))
	}
	user := userAny.(*userEntity.User)

	accessToken := a.service.IssueAccessToken(user)
	refreshToken := a.service.IssueRefreshToken(user)

	err := a.service.SaveRefreshToken(ctx, user.ID, *user.RefreshToken)

	if err != nil {
		return httpx.Redirect{}, err
	}

	state := query.String("state")
	switch state {
	case "dev":
		return httpx.Redirect{
			Location: "https://localhost:5173/auth/callback/kakao",
			Options: httpx.ResponseOptions{
				Cookies: []httpx.Cookie{
					{
						Name:     "accessToken",
						Value:    accessToken,
						Path:     "/",
						Secure:   false,
						HttpOnly: true,
						SameSite: "Lax",
						MaxAge:   int((5 * time.Minute).Seconds()),
					},
					{
						Name:     "refreshToken",
						Value:    refreshToken,
						Path:     "/",
						Secure:   false,
						HttpOnly: true,
						SameSite: "Lax",
						MaxAge:   int((7 * 24 * time.Hour).Seconds()),
					},
				},
			},
		}, nil
	case "prod":
		return httpx.Redirect{
			Location: "https://ppopgi.me/auth/callback/kakao",
			Options: httpx.ResponseOptions{
				Cookies: []httpx.Cookie{
					httpx.AccessTokenCookie(accessToken, 5*time.Minute),
					httpx.RefreshTokenCookie(refreshToken, 7*24*time.Hour),
				},
			},
		}, nil
	}
	return httpx.Redirect{}, errors.New("state가 비어있습니다.")
}

// @Summary (관리자) 어드민 계정 생성
// @Description 관리자 어드민 계정 생성 API
// @Tags Auth
// @Param req body dto.AdminUserRequest true "어드민 생성 요청"
// @Router /auth/create-admin-user [POST]
func (a *AuthController) CreateAdminUser(ctx context.Context, dto *dto.AdminUserRequest) error {
	return a.service.CreateAdminUser(ctx, dto)
}

// @Summary Access Token 재발급
// @Description Refresh Token으로 Access Token 재발급을 진행합니다.
// @Tags Auth
// @Router /auth/refresh [POST]
func (a *AuthController) RefreshToken(ctx context.Context, spineCtx spine.Ctx) (httpx.Response[string], error) {
	accessToken, ok := spineCtx.Get("auth.newAccessToken")
	if !ok {
		panic(httperr.Unauthorized("리프레시 토큰이 컨텍스트에 존재하지 않습니다."))
	}
	refreshToken, ok := spineCtx.Get("auth.newRefreshToken")
	if !ok {
		panic(httperr.Unauthorized("액세스 토큰이 컨텍스트에 존재하지 않습니다."))
	}

	return httpx.Response[string]{
		Body: "OK",
		Options: httpx.ResponseOptions{
			Cookies: []httpx.Cookie{
				httpx.AccessTokenCookie(accessToken.(string), 5*time.Minute),
				httpx.RefreshTokenCookie(refreshToken.(string), 7*24*time.Hour),
			},
		},
	}, nil
}
