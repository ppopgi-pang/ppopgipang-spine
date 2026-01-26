package controller

import "github.com/ppopgi-pang/ppopgipang-spine/moderation/service"

type ModerationController struct {
	service *service.ModerationService
}

func NewModerationController(moderationService *service.ModerationService) *ModerationController {
	return &ModerationController{service: moderationService}
}
