package service

import "github.com/ppopgi-pang/ppopgipang-spine/careers/repository"

type CareerService struct {
	repo *repository.CareerRepository
}

func NewCareerService(repo *repository.CareerRepository) *CareerService {
	return &CareerService{repo: repo}
}
