package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/utils"
	"github.com/ppopgi-pang/ppopgipang-spine/users/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/users/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(usersService *service.UserService) *UserController {
	return &UserController{service: usersService}
}

// @Summary (사용자) 내 정보 조회
// @Description 액세스 토큰으로 현재 로그인한 사용자의 정보를 조회합니다.
// @Tags Users
// @Produce json
// @Success 200 {object} dto.UserMeResponse
// @Router /users/me [GET]
func (u *UserController) GetUserInfo(ctx context.Context, spineCtx spine.Ctx) (httpx.Response[dto.UserMeResponse], error) {
	userID, err := utils.GetAuthUserID(spineCtx)
	if err != nil {
		return httpx.Response[dto.UserMeResponse]{}, err
	}
	if userID == nil {
		return httpx.Response[dto.UserMeResponse]{}, httperr.Unauthorized("인증되지 않았습니다.")
	}

	userInfo, err := u.service.GetUserInfo(ctx, *userID)
	if err != nil {
		return httpx.Response[dto.UserMeResponse]{}, err
	}

	return httpx.Response[dto.UserMeResponse]{
		Body: dto.NewUserMeResponse(userInfo),
	}, nil
}
