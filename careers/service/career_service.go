package service

import (
	"context"

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
