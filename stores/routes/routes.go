package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	authInterceptor "github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/controller"
)

func RegisterStoreRoutes(app spine.App) {
	app.Route("GET", "/stores/nearby", (*controller.StoreController).FindNearByStores, route.WithInterceptors((*authInterceptor.OptionalAccessTokenInterceptor)(nil)))
	app.Route("GET", "/stores/in-bounds", (*controller.StoreController).FindStoresInBounds, route.WithInterceptors((*authInterceptor.OptionalAccessTokenInterceptor)(nil)))
	app.Route("GET", "/stores/search", (*controller.StoreController).SearchStore)
	app.Route("GET", "/stores/summary/:storeId", (*controller.StoreController).FindByStoreSummaryId)
	app.Route("GET", "/stores/details/:storeId", (*controller.StoreController).FindByStoreDetailId, route.WithInterceptors((*authInterceptor.OptionalAccessTokenInterceptor)(nil)))
	app.Route("GET", "/stores/visits/:storeId", (*controller.StoreController).GetStoreStatById, route.WithInterceptors((*authInterceptor.OptionalAccessTokenInterceptor)(nil)))
	app.Route("GET", "/stores/reviews/:storeId", (*controller.StoreController).GetStoreReviewsById)
}
