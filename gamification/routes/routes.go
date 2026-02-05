package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/controller"
)

func RegisterGamificationRoutes(app spine.App) {
	app.Route("GET", "/v1/gamification/collections/passport", (*controller.GamificationController).GetPassport)
	app.Route("GET", "/gamification/achievement/main", (*controller.GamificationController).GetMainAchievementSummary, route.WithInterceptors((*interceptor.AccessTokenInterceptor)(nil)))
}
