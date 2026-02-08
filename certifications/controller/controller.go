package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/utils"
	"github.com/ppopgi-pang/ppopgipang-spine/certifications/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/certifications/service"
)

type CertificationController struct {
	service *service.CertificationService
}

func NewCertificationController(certificationsService *service.CertificationService) *CertificationController {
	return &CertificationController{service: certificationsService}
}

// @Summary (사용자) 방문 인증 생성
// @Description 방문 인증을 생성하고 보상(EXP, 레벨, 스탬프, 배지)을 반환합니다.
// @Tags Certifications
// @Param req body dto.CreateCheckInRequest true "방문 생성 요청 DTO"
// @Success 200 {object} dto.CertificationResponse
// @Router /certifications/checkin [POST]
func (c *CertificationController) CreateCheckin(ctx context.Context, input *dto.CreateCheckInRequest, spineCtx spine.Ctx) (httpx.Response[dto.CertificationResponse], error) {
	userId, _ := utils.GetAuthUserID(spineCtx)
	if userId == nil {
		return httpx.Response[dto.CertificationResponse]{}, httperr.Unauthorized("UserID가 존재하지 않습니다.")
	}
	result, err := c.service.CreateCheckInCertification(ctx, *userId, input)
	if err != nil {
		return httpx.Response[dto.CertificationResponse]{}, err
	}
	return httpx.Response[dto.CertificationResponse]{
		Body: result,
	}, nil
}
