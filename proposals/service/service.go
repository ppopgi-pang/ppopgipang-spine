package service

import "gorm.io/gorm"

type ProposalService struct {
	db *gorm.DB
}

func NewProposalService(db *gorm.DB) *ProposalService {
	return &ProposalService{db: db}
}
