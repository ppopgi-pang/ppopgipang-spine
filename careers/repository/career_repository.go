package repository

import "gorm.io/gorm"

type CareerRepository struct {
	db *gorm.DB
}

func NewCareerRepository(db *gorm.DB) *CareerRepository {
	return &CareerRepository{db: db}
}
