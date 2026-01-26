package service

import "github.com/ppopgi-pang/ppopgipang-spine/users/repository"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
