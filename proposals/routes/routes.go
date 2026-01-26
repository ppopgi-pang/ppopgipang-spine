package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/proposals/controller"
)

func RegisterProposalRoutes(app spine.App) {
	app.Route("POST", "/v1/proposals", (*controller.ProposalController).CreateProposal)
}
