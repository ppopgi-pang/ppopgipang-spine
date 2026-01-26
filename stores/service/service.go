package service

import "github.com/ppopgi-pang/ppopgipang-spine/stores/repository"

type StoreService struct {
	repo *repository.StoreRepository
}

func NewStoreService(repo *repository.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}
