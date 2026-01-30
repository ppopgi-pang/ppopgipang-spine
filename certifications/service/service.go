package service

import "gorm.io/gorm"

type CertificationService struct {
	db *gorm.DB
}

func NewCertificationService(db *gorm.DB) *CertificationService {
	return &CertificationService{db: db}
}
