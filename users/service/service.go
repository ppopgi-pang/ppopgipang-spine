package service

import (
	"context"

	"github.com/ppopgi-pang/ppopgipang-spine/users/entities"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) Create(user *entities.User) error {
	return u.db.Create(&user).Error
}

func (u *UserService) FindByEmail(email string) (entities.User, error) {
	var user entities.User

	err := u.db.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (u *UserService) GetUserInfo(ctx context.Context, userID int64) (entities.User, error) {
	var user entities.User

	err := u.db.WithContext(ctx).
		Where("id = ?", userID).
		First(&user).
		Error
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
