package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/reviews/controller"
)

func RegisterReviewRoutes(app spine.App) {
	app.Route("GET", "/v1/reviews", (*controller.ReviewController).GetReviews)
}
