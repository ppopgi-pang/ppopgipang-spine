package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	authInterceptor "github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/controller"
)

func RegisterUserRoutes(app spine.App) {
	app.Route("GET", "/careers/job-postings", (*controller.CareerController).GetJobPostings)
	app.Route("POST", "/careers/job-postings", (*controller.CareerController).CreateJobPosting, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
	app.Route("GET", "/careers/job-postings/:id", (*controller.CareerController).GetJobPosting)
	app.Route("PATCH", "/careers/job-postings/:id", (*controller.CareerController).UpdateJobPosting, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
	app.Route("DELETE", "/careers/job-postings/:id", (*controller.CareerController).DeleteJobPosting, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
	app.Route("POST", "/careers/applications", (*controller.CareerController).CreateApplication)
	app.Route("GET", "/careers/applications", (*controller.CareerController).GetApplications)
	app.Route("GET", "/careers/applications/:id", (*controller.CareerController).GetApplication, route.WithInterceptors((*authInterceptor.AccessTokenInterceptor)(nil)))
}
