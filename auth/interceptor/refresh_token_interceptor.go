package interceptor

import (
	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/service"
)

type RefreshTokenInterceptor struct {
	auth *service.AuthService
}

type RefreshPayload struct {
	UserID int64
	Role   string
}

func NewRefreshTokenInterceptor(auth *service.AuthService) *RefreshTokenInterceptor {
	return &RefreshTokenInterceptor{
		auth: auth,
	}
}

func (i *RefreshTokenInterceptor) PreHandle(ctx core.ExecutionContext, meta core.HandlerMeta) error {
	refreshToken := extractCookieToken(ctx, "refreshToken")
	if refreshToken == "" {
		return httperr.Unauthorized("refresh token이 없습니다.")
	}

	result, err := i.auth.RotateRefreshToken(refreshToken)
	if err != nil {
		return httperr.Unauthorized("유효하지 않은 refresh token")
	}

	ctx.Set("auth.userId", result.Payload.UserID)
	ctx.Set("auth.tokenId", result.Payload.TokenID)
	ctx.Set("auth.tokenType", "refresh")

	ctx.Set("auth.newRefreshToken", result.NewRefreshToken)

	accessToken := i.auth.IssueAccessToken(&result.User)

	ctx.Set("auth.newAccessToken", accessToken)

	return nil
}

func (i *RefreshTokenInterceptor) PostHandle(ctx core.ExecutionContext, meta core.HandlerMeta) {
}

func (i *RefreshTokenInterceptor) AfterCompletion(ctx core.ExecutionContext, meta core.HandlerMeta, err error) {
}
