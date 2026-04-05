package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	authInterceptor "github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	"github.com/ppopgi-pang/ppopgipang-spine/users/controller"
)

func RegisterUserRoutes(app spine.App) {
	app.Route("GET", "/users/me", (*controller.UserController).GetUserInfo, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
}
