package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/ppopgi-pang/ppopgipang-spine/reviews/service"
)

type ReviewController struct {
	service *service.ReviewService
}

func NewReviewController(reviewsService *service.ReviewService) *ReviewController {
	return &ReviewController{service: reviewsService}
}

func (r *ReviewController) GetReviews(ctx context.Context, query query.Values) {

}
