package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/service"
)

type CareerController struct {
	service *service.CareerService
}

func NewCareerController(careerService *service.CareerService) *CareerController {
	return &CareerController{service: careerService}
}

func (c *CareerController) GetJobPostings(ctx context.Context, query query.Values) {

}
