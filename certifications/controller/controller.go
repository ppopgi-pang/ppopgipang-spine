package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/ppopgi-pang/ppopgipang-spine/certifications/service"
)

type CertificationController struct {
	service *service.CertificationService
}

func NewCertificationController(certificationsService *service.CertificationService) *CertificationController {
	return &CertificationController{service: certificationsService}
}

func (c *CertificationController) GetLootGallery(ctx context.Context, query query.Values) {

}
