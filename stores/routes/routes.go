package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	authInterceptor "github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/controller"
)

func RegisterStoreRoutes(app spine.App) {
	app.Route("GET", "/stores/nearby", (*controller.StoreController).FindNearByStores, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
	app.Route("GET", "/stores/in-bounds", (*controller.StoreController).FindStoresInBounds, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
}
