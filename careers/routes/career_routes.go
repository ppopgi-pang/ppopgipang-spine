package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/controller"
)

func RegisterUserRoutes(app spine.App) {
	app.Route("POST", "v1/careers/job-postings", (*controller.CareerController).GetJobPostings)
}
