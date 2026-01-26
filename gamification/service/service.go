package service

import "github.com/ppopgi-pang/ppopgipang-spine/gamification/repository"

type GamificationService struct {
	repo *repository.GamificationRepository
}

func NewGamificationService(repo *repository.GamificationRepository) *GamificationService {
	return &GamificationService{repo: repo}
}
