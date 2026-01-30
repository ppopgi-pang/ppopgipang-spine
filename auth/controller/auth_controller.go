package controller

import (
	"context"
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

func (a *AuthController) KakaoCallback(ctx context.Context, code query.Values, spineCtx spine.Ctx) httpx.Redirect {
	userAny, ok := spineCtx.Get("auth.user")
	if !ok {
		panic(httperr.Unauthorized("인증 정보가 없습니다."))
	}
	user := userAny.(*userEntity.User)

	accessToken := a.service.IssueAccessToken(user)
	refreshToken := a.service.IssueRefreshToken(user)

	return httpx.Redirect{
		Location: "http://localhost:3000/auth/kakao/callback",
		Options: httpx.ResponseOptions{
			Cookies: []httpx.Cookie{
				httpx.AccessTokenCookie(accessToken, 5*time.Minute),
				httpx.RefreshTokenCookie(refreshToken, 7*24*time.Hour),
			},
		},
	}
}

// @Summary (관리자) 어드민 계정 생성
// @Description 관리자 어드민 계정 생성 API
// @Tags Auth
// @Param req body dto.AdminUserRequest true "어드민 생성 요청"
// @Router /auth/create-admin-user [POST]
func (a *AuthController) CreateAdminUser(ctx context.Context, dto *dto.AdminUserRequest) error {
	return a.service.CreateAdminUser(ctx, dto)
}
