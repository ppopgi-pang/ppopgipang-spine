package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/interceptor"
	"github.com/ppopgi-pang/ppopgipang-spine/certifications/controller"
)

func RegisterCertificationRoutes(app spine.App) {
	app.Route("POST", "/certifications/checkin", (*controller.CertificationController).CreateCheckin, route.WithInterceptors((*interceptor.AccessTokenInterceptor)(nil)))
}
