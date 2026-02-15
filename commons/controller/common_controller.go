package controller

import (
	"context"
	"errors"

	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/multipart"
	"github.com/NARUBROWN/spine/pkg/path"
	"github.com/ppopgi-pang/ppopgipang-spine/commons/service"
)

type CommonController struct {
	commonService *service.CommonService
}

func NewCommonController(commonService *service.CommonService) *CommonController {
	return &CommonController{commonService: commonService}
}

// @Summary (공통) 파일 업로드
// @Description 파일 업로드 API
// @Tags Commons
// @Accept multipart/form-data
// @Param files formData file true "업로드할 파일들"
// @Success 200 {array} string
// @Router /commons/file-uploads [POST]
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

// @Summary (공통) 이미지 확인
// @Description 이미지 확인 API
// @Tags Commons
// @Param path path string true "파일 경로"
// @Param fileName path string true "파일명 (확장자 포함)"
// @Router /commons/images/{path}/{fileName} [GET]
func (c *CommonController) GetFile(ctx context.Context, path path.String, fileName path.String) (httpx.Binary, error) {
	result, err := c.commonService.ReadFromDisk(path.Value, fileName.Value)
	if err != nil {
		return httpx.Binary{}, err
	}
	return httpx.Binary{
		ContentType: "image/png",
		Data:        result,
	}, nil
}
