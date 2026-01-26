package service

import "github.com/ppopgi-pang/ppopgipang-spine/proposals/repository"

type ProposalService struct {
	repo *repository.ProposalRepository
}

func NewProposalService(repo *repository.ProposalRepository) *ProposalService {
	return &ProposalService{repo: repo}
}
