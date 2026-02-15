package service

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

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

func (c *CommonService) ReadFromDisk(path string, fileName string) ([]byte, error) {
	if strings.TrimSpace(path) == "" || strings.TrimSpace(fileName) == "" {
		return nil, errors.New("경로와 파일명이 필요합니다")
	}

	if filepath.IsAbs(path) {
		return nil, errors.New("유효하지 않은 경로입니다")
	}

	if strings.Contains(path, "..") || strings.Contains(fileName, "..") {
		return nil, errors.New("유효하지 않은 경로입니다")
	}

	file := filepath.Base(fileName)
	if strings.Contains(file, string(filepath.Separator)) {
		return nil, errors.New("유효하지 않은 파일명입니다")
	}

	baseDir, err := filepath.Abs(filepath.Join(".", "uploads"))
	if err != nil {
		return nil, err
	}

	targetPath := filepath.Join(baseDir, path, file)
	if filepath.Clean(targetPath) != targetPath {
		return nil, errors.New("유효하지 않은 경로입니다")
	}

	targetPath, err = filepath.Abs(targetPath)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(targetPath, baseDir+string(filepath.Separator)) {
		return nil, errors.New("유효하지 않은 경로입니다")
	}

	data, err := os.ReadFile(targetPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
