package controller

import (
	"context"

	"github.com/ppopgi-pang/ppopgipang-spine/users/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(usersService *service.UserService) *UserController {
	return &UserController{service: usersService}
}

func (u *UserController) GetUserInfo(ctx context.Context) {}
