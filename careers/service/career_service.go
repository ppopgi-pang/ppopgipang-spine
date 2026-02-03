package service

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ppopgi-pang/ppopgipang-spine/careers/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/entities"
	"gorm.io/gorm"
)

type CareerService struct {
	db *gorm.DB
}

func NewCareerService(db *gorm.DB) *CareerService {
	return &CareerService{db: db}
}

func (c *CareerService) GetJobPostings(ctx context.Context, isActive bool, page, size int) (dto.JobPostingListResponse, error) {
	db := c.db.WithContext(ctx)

	type jobPostingWithCount struct {
		entities.JobPosting
		ApplicationCount int64 `gorm:"column:applicationCount"`
	}

	var rows []jobPostingWithCount
	var total int64

	query := db.
		Table("job_postings AS jobPosting").
		Select(`
			jobPosting.*,
			COUNT(application.id) AS applicationCount
		`).
		Joins(`
			LEFT JOIN applications as application
			ON application.jobPosting_id = jobPosting_id
		`).
		Group("jobPosting.id").
		Order("jobPosting.createdAt DESC").
		Offset((page - 1) * size).
		Limit(size)

	if isActive {
		query = query.Where("jobPosting.isActive = ?", isActive)
	}

	if err := query.Scan(&rows).Error; err != nil {
		return dto.JobPostingListResponse{}, err
	}

	countQuery := db.Model(entities.JobPosting{})
	if isActive {
		countQuery = countQuery.Where("isActive = ?", isActive)
	}

	if err := countQuery.Count(&total).Error; err != nil {
		return dto.JobPostingListResponse{}, err
	}

	items := make([]dto.JobPostingResponse, 0, len(rows))

	for _, row := range rows {
		desc := ""
		if row.Description != nil {
			desc = *row.Description
		}

		dept := ""
		if row.Department != nil {
			dept = *row.Department
		}

		posType := ""
		if row.PositionType != nil {
			posType = *row.PositionType
		}

		loc := ""
		if row.Location != nil {
			loc = *row.Location
		}

		items = append(items, dto.JobPostingResponse{
			ID:               row.ID,
			Title:            row.Title,
			Description:      desc,
			Department:       dept,
			PositionType:     posType,
			Location:         loc,
			IsActive:         row.IsActive,
			CreatedAt:        row.CreatedAt,
			UpdatedAt:        row.UpdatedAt,
			ApplicationCount: row.ApplicationCount,
		})
	}

	return dto.JobPostingListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (c *CareerService) CreateJobPosting(ctx context.Context, dto *dto.JobPostingRequest) error {
	db := c.db.WithContext(ctx)

	job := entities.JobPosting{
		Title:        dto.Title,
		Description:  &dto.Description,
		Department:   &dto.Department,
		PositionType: &dto.PositionType,
		Location:     &dto.Location,
		IsActive:     true,
	}

	if err := db.Create(&job).Error; err != nil {
		return err
	}

	return nil
}

func (c *CareerService) GetJobPosting(ctx context.Context, id int64) (dto.JobPostingResponse, error) {
	db := c.db.WithContext(ctx)

	var job entities.JobPosting

	err := db.First(&job, id).Error
	if err != nil {
		return dto.JobPostingResponse{}, err
	}

	count := db.Model(&job).Association("Applications").Count()

	des := ""
	if job.Description != nil {
		des = *job.Description
	}

	dept := ""
	if job.Department != nil {
		dept = *job.Department
	}

	positionType := ""
	if job.PositionType != nil {
		positionType = *job.PositionType
	}

	location := ""
	if job.Location != nil {
		location = *job.Location
	}

	return dto.JobPostingResponse{
		ID:               job.ID,
		Title:            job.Title,
		Description:      des,
		Department:       dept,
		PositionType:     positionType,
		Location:         location,
		IsActive:         job.IsActive,
		CreatedAt:        job.CreatedAt,
		UpdatedAt:        job.UpdatedAt,
		ApplicationCount: count,
	}, nil
}

func (c *CareerService) UpdateJobPosting(ctx context.Context, id int64, dto *dto.JobPostingModifyRequest) error {
	db := c.db.WithContext(ctx)

	var job entities.JobPosting
	if err := db.First(&job, id).Error; err != nil {
		return err
	}

	if dto.Title != nil {
		job.Title = *dto.Title
	}
	if dto.Description != nil {
		job.Description = dto.Description
	}
	if dto.Department != nil {
		job.Department = dto.Department
	}
	if dto.PositionType != nil {
		job.PositionType = dto.PositionType
	}
	if dto.Location != nil {
		job.Location = dto.Location
	}
	if dto.IsActive != nil {
		job.IsActive = *dto.IsActive
	}

	return db.Save(&job).Error
}

func (c *CareerService) DeleteJobPosting(ctx context.Context, id int64) error {
	db := c.db.WithContext(ctx)

	return db.Delete(&entities.JobPosting{}, id).Error
}

func (c *CareerService) CreateApplication(ctx context.Context, dto *dto.CreateApplicationRequest) (int64, error) {
	db := c.db.WithContext(ctx)

	if dto.ResumeFileName != nil && *dto.ResumeFileName != "" {
		if err := moveResumeFromTemp(*dto.ResumeFileName); err != nil {
			return 0, err
		}
	}

	application := &entities.Application{
		JobPostingID: dto.JobPostingId,
		Name:         dto.Name,
		Email:        dto.Email,
		Phone:        dto.Phone,
		ResumeName:   dto.ResumeFileName,
		Memo:         dto.Memo,
		Status:       "new",
	}

	if err := db.Create(&application).Error; err != nil {
		return 0, err
	}

	return application.ID, nil
}

func (c *CareerService) GetApplications(ctx context.Context, jobPostingId int64, status string, page, size int) (dto.ApplicationListResponse, error) {
	db := c.db.WithContext(ctx)

	query := db.Model(&entities.Application{})

	if jobPostingId != 0 {
		query = query.Where("jobPosting_id = ?", jobPostingId)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return dto.ApplicationListResponse{}, err
	}

	var applications []entities.Application
	if err := query.
		Preload("JobPosting").
		Order("createdAt DESC").
		Offset((page - 1) * size).
		Limit(size).
		Find(&applications).Error; err != nil {
		return dto.ApplicationListResponse{}, err
	}

	items := make([]dto.ApplicationResponse, 0, len(applications))
	for _, row := range applications {
		items = append(items, dto.ApplicationResponse{
			ID: row.ID,
			JobPostingResponse: dto.ApplicationJobPosting{
				ID:    row.JobPosting.ID,
				Title: row.JobPosting.Title,
			},
			Name:           row.Name,
			Email:          row.Email,
			Phone:          row.Phone,
			ResumeFileName: row.ResumeName,
			Memo:           row.Memo,
			Status:         row.Status,
			CreatedAt:      row.CreatedAt,
		})
	}

	return dto.ApplicationListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (c *CareerService) GetApplication(ctx context.Context, id int64) (dto.ApplicationResponse, error) {
	db := c.db.WithContext(ctx)

	var application entities.Application
	err := db.
		Preload("JobPosting").
		First(&application, id).
		Error

	if err != nil {
		return dto.ApplicationResponse{}, err
	}

	return dto.ApplicationResponse{
		ID: application.ID,
		JobPostingResponse: dto.ApplicationJobPosting{
			ID:    application.JobPostingID,
			Title: application.JobPosting.Title,
		},
		Name:           application.Name,
		Email:          application.Email,
		Phone:          application.Phone,
		ResumeFileName: application.ResumeName,
		Memo:           application.Memo,
		Status:         application.Status,
		CreatedAt:      application.CreatedAt,
	}, nil
}

func moveResumeFromTemp(filename string) error {
	tempDir := filepath.Join("uploads", "temps")
	careerDir := filepath.Join("uploads", "careers")

	if err := os.MkdirAll(careerDir, 0755); err != nil {
		return err
	}

	oldPath := filepath.Join(tempDir, filename)
	newPath := filepath.Join(careerDir, filename)

	if err := os.Rename(oldPath, newPath); err != nil {
		if errors.Is(err, syscall.EXDEV) {
			if err := copyFile(oldPath, newPath); err != nil {
				return err
			}
			return os.Remove(oldPath)
		}
		return err
	}

	return nil
}

func copyFile(srcPath, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
