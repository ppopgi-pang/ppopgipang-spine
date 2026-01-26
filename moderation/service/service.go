package service

import "github.com/ppopgi-pang/ppopgipang-spine/moderation/repository"

type ModerationService struct {
	repo *repository.ModerationRepository
}

func NewModerationService(repo *repository.ModerationRepository) *ModerationService {
	return &ModerationService{repo: repo}
}
