package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/controller"
)

func RegisterGamificationRoutes(app spine.App) {
	app.Route("GET", "/v1/gamification/collections/passport", (*controller.GamificationController).GetPassport)
}
