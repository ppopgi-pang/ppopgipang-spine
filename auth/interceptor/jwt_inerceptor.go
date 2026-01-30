package interceptor

import (
	"strings"

	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/service"
)

type JwtInterceptor struct {
	auth *service.AuthService
}

func NewJwtInterceptor(auth *service.AuthService) *JwtInterceptor {
	return &JwtInterceptor{auth: auth}
}

func (i *JwtInterceptor) PreHandle(ctx core.ExecutionContext, meta core.HandlerMeta) error {
	accessToken := extractCookieToken(ctx, "accessToken")
	if accessToken != "" {
		payload, err := i.auth.VerifyAccessToken(accessToken)
		if err != nil {
			return httperr.Unauthorized("유효하지 않은 access token")
		}

		ctx.Set("auth.userId", payload.UserID)
		ctx.Set("auth.role", payload.Role)
		ctx.Set("auth.tokenType", "access")
		return nil
	}

	refreshToken := extractCookieToken(ctx, "refreshToken")
	if refreshToken == "" {
		return httperr.Unauthorized("token이 없습니다.")
	}

	payload, err := i.auth.VerifyRefreshToken(refreshToken)
	if err != nil {
		return httperr.Unauthorized("유효하지 않은 refresh token")
	}

	ctx.Set("auth.userId", payload.UserID)
	ctx.Set("auth.tokenId", payload.TokenID)
	ctx.Set("auth.tokenType", "refresh")
	return nil
}

func (i *JwtInterceptor) PostHandle(ctx core.ExecutionContext, meta core.HandlerMeta) {
}

func (i *JwtInterceptor) AfterCompletion(ctx core.ExecutionContext, meta core.HandlerMeta, err error) {
}

func extractCookieToken(ctx core.ExecutionContext, key string) string {
	return extractCookie(ctx.Header("Cookie"), key)
}

func extractCookie(cookieHeader, key string) string {
	if cookieHeader == "" {
		return ""
	}

	parts := strings.SplitSeq(cookieHeader, ";")
	for part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 && kv[0] == key {
			return kv[1]
		}
	}

	return ""
}
