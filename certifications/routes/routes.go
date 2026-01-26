package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/certifications/controller"
)

func RegisterCertificationRoutes(app spine.App) {
	app.Route("GET", "/v1/certifications/loots", (*controller.CertificationController).GetLootGallery)
}
