package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/controller"
)

func RegisterUserRoutes(app spine.App) {
	app.Route("GET", "/careers/job-postings", (*controller.CareerController).GetJobPostings)
	app.Route("POST", "/careers/job-postings", (*controller.CareerController).CreateJobPosting)
	app.Route("GET", "/careers/job-postings/:id", (*controller.CareerController).GetJobPosting)
}
