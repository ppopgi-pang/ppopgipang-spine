package interceptor

import (
	"strings"

	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/service"
)

type AccessTokenInterceptor struct {
	auth *service.AuthService
}

func NewAccessTokenInterceptor(auth *service.AuthService) *AccessTokenInterceptor {
	return &AccessTokenInterceptor{auth: auth}
}

func (i *AccessTokenInterceptor) PreHandle(ctx core.ExecutionContext, meta core.HandlerMeta) error {
	accessToken := extractCookieToken(ctx, "accessToken")
	if accessToken == "" {
		return httperr.Unauthorized("access token이 없습니다.")
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

func (i *AccessTokenInterceptor) PostHandle(ctx core.ExecutionContext, meta core.HandlerMeta) {
}

func (i *AccessTokenInterceptor) AfterCompletion(ctx core.ExecutionContext, meta core.HandlerMeta, err error) {
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
