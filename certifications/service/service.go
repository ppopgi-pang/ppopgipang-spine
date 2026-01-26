package service

import "github.com/ppopgi-pang/ppopgipang-spine/certifications/repository"

type CertificationService struct {
	repo *repository.CertificationRepository
}

func NewCertificationService(repo *repository.CertificationRepository) *CertificationService {
	return &CertificationService{repo: repo}
}
