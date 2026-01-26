package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/service"
)

type GamificationController struct {
	service *service.GamificationService
}

func NewGamificationController(gamificationService *service.GamificationService) *GamificationController {
	return &GamificationController{service: gamificationService}
}

func (g *GamificationController) GetPassport(ctx context.Context, query query.Values) {

}
