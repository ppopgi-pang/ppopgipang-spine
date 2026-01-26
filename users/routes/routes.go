package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/users/controller"
)

func RegisterUserRoutes(app spine.App) {
	app.Route("GET", "/v1/users", (*controller.UserController).GetUserInfo)
}
