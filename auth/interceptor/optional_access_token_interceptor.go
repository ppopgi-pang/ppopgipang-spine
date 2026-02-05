package interceptor

import (
	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/service"
)

type OptionalAccessTokenInterceptor struct {
	auth *service.AuthService
}

func NewOptionalAccessTokenInterceptor(auth *service.AuthService) *OptionalAccessTokenInterceptor {
	return &OptionalAccessTokenInterceptor{auth: auth}
}

func (i *OptionalAccessTokenInterceptor) PreHandle(ctx core.ExecutionContext, meta core.HandlerMeta) error {
	accessToken := extractCookieToken(ctx, "accessToken")
	if accessToken == "" {
		ctx.Set("auth.userId", "")
		ctx.Set("auth.role", "")
		ctx.Set("auth.tokenType", "")
		return nil
	}

	payload, err := i.auth.VerifyAccessToken(accessToken)
	if err != nil {
		return httperr.Unauthorized("유효하지 않은 access token")
	}

	ctx.Set("auth.userId", payload.UserID)
	ctx.Set("auth.role", payload.Role)
	ctx.Set("auth.tokenType", "access")
	return nil
}

func (i *OptionalAccessTokenInterceptor) PostHandle(ctx core.ExecutionContext, meta core.HandlerMeta) {
}

func (i *OptionalAccessTokenInterceptor) AfterCompletion(ctx core.ExecutionContext, meta core.HandlerMeta, err error) {
}
