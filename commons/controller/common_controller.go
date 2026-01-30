package controller

import (
	"context"
	"errors"

	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/multipart"
	"github.com/ppopgi-pang/ppopgipang-spine/commons/service"
)

type CommonController struct {
	commonService *service.CommonService
}

func NewCommonController(commonService *service.CommonService) *CommonController {
	return &CommonController{commonService: commonService}
}

func (c *CommonController) UploadFiles(ctx context.Context, files multipart.UploadedFiles) (httpx.Response[[]string], error) {
	names := make([]string, 0, len(files.Files))
	for _, file := range files.Files {

		if file.Size > 10<<20 {
			return httpx.Response[[]string]{}, errors.New("파일이 너무 큽니다.")
		}

		name, err := c.commonService.SaveToDisk(file, "./uploads/temps")
		if err != nil {
			return httpx.Response[[]string]{}, err
		}

		names = append(names, name)
	}

	return httpx.Response[[]string]{
		Body: names,
	}, nil
}
