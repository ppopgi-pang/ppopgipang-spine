package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/commons/controller"
)

func RegisterCommonRoutes(app spine.App) {
	app.Route("POST", "/commons/file-uploads", (*controller.CommonController).UploadFiles)
	app.Route("GET", "/commons/images/:path/:fileName", (*controller.CommonController).GetFile)
}
