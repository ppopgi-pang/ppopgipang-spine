package repository

import (
	"github.com/ppopgi-pang/ppopgipang-spine/users/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user *entities.User) error {
	return u.db.Create(&user).Error
}

func (u UserRepository) FindByEmail(email string) (entities.User, error) {
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
