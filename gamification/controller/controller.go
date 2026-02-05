package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/utils"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/dto"
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

// @Summary (사용자) 가장 최근의 미해결 업적 진행상황 조회
// @Description 맵 화면에서 가장 최근의 미해결된 업적의 진행상황을 조회하는 API입니다.
// @Tags Gamification
// @Success 200 {object} dto.GamificationMainAchievementResponse
// @Router /gamification/achievement/main [GET]
func (g *GamificationController) GetMainAchievementSummary(ctx context.Context, spineCtx spine.Ctx) (httpx.Response[dto.GamificationMainAchievementResponse], error) {
	userID, err := utils.GetAuthUserID(spineCtx)
	if err != nil {
		return httpx.Response[dto.GamificationMainAchievementResponse]{}, err
	}
	if userID == nil {
		return httpx.Response[dto.GamificationMainAchievementResponse]{}, httperr.Unauthorized("인증되지 않았습니다.")
	}

	result, err := g.service.GetMainAchievementSummary(ctx, *userID)
	if err != nil {
		return httpx.Response[dto.GamificationMainAchievementResponse]{}, err
	}

	return httpx.Response[dto.GamificationMainAchievementResponse]{
		Body: result,
	}, nil
}
