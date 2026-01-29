package service

import (
	"io"
	"os"
	"path/filepath"

	"github.com/NARUBROWN/spine/pkg/multipart"
	"github.com/google/uuid"
)

type CommonService struct{}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (c *CommonService) SaveToDisk(file multipart.UploadedFile, dir string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	name := uuid.NewString() + filepath.Ext(file.Filename)

	path := filepath.Join(dir, name)

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return name, err
}
