package repository

import "gorm.io/gorm"

type CertificationRepository struct {
	db *gorm.DB
}

func NewCertificationRepository(db *gorm.DB) *CertificationRepository {
	return &CertificationRepository{db: db}
}
