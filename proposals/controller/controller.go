package controller

import (
	"context"

	"github.com/ppopgi-pang/ppopgipang-spine/proposals/service"
)

type ProposalController struct {
	service *service.ProposalService
}

func NewProposalController(proposalsService *service.ProposalService) *ProposalController {
	return &ProposalController{service: proposalsService}
}

func (p *ProposalController) CreateProposal(ctx context.Context) {

}
