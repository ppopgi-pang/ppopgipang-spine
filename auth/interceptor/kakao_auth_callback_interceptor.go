package interceptor

import (
	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/client"
)

type KakaoAuthCallbackInterceptor struct {
	client *client.KakaoOAuthClient
}

func NewKakaoAuthCallbackInterceptor(client *client.KakaoOAuthClient) *KakaoAuthCallbackInterceptor {
	return &KakaoAuthCallbackInterceptor{client: client}
}

func (i *KakaoAuthCallbackInterceptor) PreHandle(ctx core.ExecutionContext, meta core.HandlerMeta) error {
	codes := ctx.Queries()["code"]
	if len(codes) == 0 || codes[0] == "" {
		return httperr.Unauthorized("kakao oauth code가 없습니다.")
	}
	code := codes[0]

	// 2. code → kakao access token 교환
	token, err := i.client.ExchangeCodeForToken(code)
	if err != nil {
		return httperr.Unauthorized("kakao token 발급 실패")
	}

	// 3. access token으로 kakao user 조회
	kakaoUser, err := i.client.GetUser(token)
	if err != nil {
		return httperr.Unauthorized("kakao 사용자 정보 조회 실패")
	}

	// 4. 내부 User 조회 또는 생성
	user, err := i.client.MapOrCreateUser(kakaoUser)
	if err != nil {
		return httperr.Unauthorized("사용자 처리 중 오류 발생")
	}

	// 5. Spine Context에 주입
	ctx.Set("auth.user", user)

	return nil
}

func (i *KakaoAuthCallbackInterceptor) PostHandle(ctx core.ExecutionContext, meta core.HandlerMeta) {
}

func (i *KakaoAuthCallbackInterceptor) AfterCompletion(ctx core.ExecutionContext, meta core.HandlerMeta, err error) {
}
