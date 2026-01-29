package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/controller"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
)

func RegisterAuthRoutes(app spine.App) {
	app.Route("GET", "/auth/kakao/callback", (*controller.AuthController).KakaoCallback, route.WithInterceptors((*interceptor.KakaoAuthCallbackInterceptor)(nil)))
}
