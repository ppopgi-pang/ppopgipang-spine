package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/controller"
)

func RegisterStoreRoutes(app spine.App) {
	app.Route("GET", "/v1/stores/nearest", (*controller.StoreController).FindNearestStore)
}
