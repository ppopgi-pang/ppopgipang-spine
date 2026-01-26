package service

import "github.com/ppopgi-pang/ppopgipang-spine/reviews/repository"

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}
