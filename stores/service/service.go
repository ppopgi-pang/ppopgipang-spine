package service

import (
	"context"
	"math"

	certEntity "github.com/ppopgi-pang/ppopgipang-spine/certifications/entities"
	reviewEntity "github.com/ppopgi-pang/ppopgipang-spine/reviews/entities"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	"gorm.io/gorm"
)

type StoreService struct {
	db *gorm.DB
}

func NewStoreService(db *gorm.DB) *StoreService {
	return &StoreService{db: db}
}

func (s *StoreService) FindNearByStores(ctx context.Context, latitude, longitude float64, radius int64, page, size int, keyword, filter string, userID *int64) (dto.FindNearByDto, error) {
	type storeWithDistance struct {
		entities.Store `gorm:"embedded"`
		Distance       float64 `gorm:"column:distance"`
		ReviewCount    int     `gorm:"column:review_count"`
		RecentReview   *string `gorm:"column:recent_review"`
		RecentCertCnt  int     `gorm:"column:recent_cert_count"`
	}

	if filter == "scrapped" && userID == nil {
		return dto.FindNearByDto{
			Success: true,
			Data:    []dto.StoreFindNearByResponse{},
			Meta: dto.Meta{
				Count: 0,
			},
		}, nil
	}

	baseQuery := s.db.WithContext(ctx).
		Model(&entities.Store{}).
		Where(
			"ST_Distance_Sphere(POINT(?, ?), POINT(stores.longitude, stores.latitude)) <= ?",
			longitude,
			latitude,
			radius,
		)

	if keyword != "" {
		baseQuery = baseQuery.Where("stores.name LIKE ?", "%"+keyword+"%")
	}

	if filter == "scrapped" {
		baseQuery = baseQuery.Joins(
			"JOIN "+certEntity.UserStoreStat{}.TableName()+" uss ON uss.storeId = stores.id AND uss.userId = ? AND uss.isScrapped = 1",
			*userID,
		)
	}

	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return dto.FindNearByDto{}, err
	}

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	reviewCountSub := s.db.
		Table(reviewEntity.Review{}.TableName()).
		Select("storeId, COUNT(*) as review_count").
		Group("storeId")

	recentReviewSub := s.db.
		Table(reviewEntity.Review{}.TableName()).
		Select("storeId, content, ROW_NUMBER() OVER (PARTITION BY storeId ORDER BY createdAt DESC, id DESC) as rn")

	recentCertCountSub := s.db.
		Table(certEntity.Certification{}.TableName()).
		Select("storeId, COUNT(*) as recent_cert_count").
		Where("occurredAt >= DATE_SUB(NOW(), INTERVAL 3 HOUR)").
		Group("storeId")

	var rows []storeWithDistance
	queryBuilder := baseQuery.
		Session(&gorm.Session{}).
		Select(
			"stores.*, "+
				"ST_Distance_Sphere(POINT(?, ?), POINT(stores.longitude, stores.latitude)) as distance, "+
				"rc.review_count as review_count, "+
				"rr.content as recent_review, "+
				"COALESCE(rcc.recent_cert_count, 0) as recent_cert_count",
			longitude,
			latitude,
		).
		Joins("LEFT JOIN (?) rc ON rc.storeId = stores.id", reviewCountSub).
		Joins("LEFT JOIN (?) rr ON rr.storeId = stores.id AND rr.rn = 1", recentReviewSub).
		Joins("LEFT JOIN (?) rcc ON rcc.storeId = stores.id", recentCertCountSub).
		Preload("Type")

	switch filter {
	case "popular":
		queryBuilder = queryBuilder.
			Order("rc.review_count DESC").
			Order("stores.averageRating DESC").
			Order("rcc.recent_cert_count DESC").
			Order("distance ASC")
	case "recent_cert":
		queryBuilder = queryBuilder.
			Order("rcc.recent_cert_count DESC").
			Order("distance ASC")
	default:
		queryBuilder = queryBuilder.Order("distance ASC")
	}

	err := queryBuilder.
		Offset((page - 1) * size).
		Limit(size).
		Find(&rows).
		Error
	if err != nil {
		return dto.FindNearByDto{}, err
	}

	data := make([]dto.StoreFindNearByResponse, 0, len(rows))
	for _, row := range rows {
		store := row.Store

		var storeType dto.StoreTypeResponse
		if store.Type != nil {
			storeType = dto.StoreTypeResponse{
				ID:          store.Type.ID,
				Name:        store.Type.Name,
				Description: store.Type.Description,
			}
		}

		data = append(data, dto.StoreFindNearByResponse{
			ID:            store.ID,
			Name:          store.Name,
			Address:       store.Address,
			Region1:       store.Region1,
			Region2:       store.Region2,
			Latitude:      store.Latitude,
			Longitude:     store.Longitude,
			Phone:         store.Phone,
			AverageRating: store.AverageRating,
			Distance:      int(math.Round(row.Distance)),
			Type:          storeType,
			RecentReview:  row.RecentReview,
			ReviewCount:   row.ReviewCount,
			CreatedAt:     store.CreatedAt,
			UpdatedAt:     store.UpdatedAt,
		})
	}

	return dto.FindNearByDto{
		Success: true,
		Data:    data,
		Meta: dto.Meta{
			Count: total,
		},
	}, nil
}

func (s *StoreService) FindStoresInBounds(ctx context.Context, north, south, east, west float64, keyword, filter string, userID *int64) (dto.FindInBoundsDto, error) {
	type storeWithReview struct {
		entities.Store `gorm:"embedded"`
		ReviewCount    int     `gorm:"column:review_count"`
		RecentReview   *string `gorm:"column:recent_review"`
		RecentCertCnt  int     `gorm:"column:recent_cert_count"`
	}

	if filter == "scrapped" && userID == nil {
		return dto.FindInBoundsDto{
			Success: true,
			Data:    []dto.StoreInBoundsResponse{},
			Meta: dto.Meta{
				Count: 0,
			},
		}, nil
	}

	baseQuery := s.db.WithContext(ctx).
		Model(&entities.Store{}).
		Where("stores.latitude BETWEEN ? AND ?", south, north).
		Where("stores.longitude BETWEEN ? AND ?", west, east)

	if keyword != "" {
		likeKeyword := "%" + keyword + "%"
		baseQuery = baseQuery.Where(
			"(stores.name LIKE ? OR stores.address LIKE ?)",
			likeKeyword,
			likeKeyword,
		)
	}

	if filter == "scrapped" {
		baseQuery = baseQuery.Joins(
			"JOIN "+certEntity.UserStoreStat{}.TableName()+" uss ON uss.storeId = stores.id AND uss.userId = ? AND uss.isScrapped = 1",
			*userID,
		)
	}

	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return dto.FindInBoundsDto{}, err
	}

	reviewCountSub := s.db.
		Table(reviewEntity.Review{}.TableName()).
		Select("storeId, COUNT(*) as review_count").
		Group("storeId")

	recentReviewSub := s.db.
		Table(reviewEntity.Review{}.TableName()).
		Select("storeId, content, ROW_NUMBER() OVER (PARTITION BY storeId ORDER BY createdAt DESC, id DESC) as rn")

	recentCertCountSub := s.db.
		Table(certEntity.Certification{}.TableName()).
		Select("storeId, COUNT(*) as recent_cert_count").
		Where("occurredAt >= DATE_SUB(NOW(), INTERVAL 3 HOUR)").
		Group("storeId")

	var stores []storeWithReview
	queryBuilder := baseQuery.
		Session(&gorm.Session{}).
		Select(
			"stores.*, "+
				"rc.review_count as review_count, "+
				"rr.content as recent_review, "+
				"COALESCE(rcc.recent_cert_count, 0) as recent_cert_count",
		).
		Joins("LEFT JOIN (?) rc ON rc.storeId = stores.id", reviewCountSub).
		Joins("LEFT JOIN (?) rr ON rr.storeId = stores.id AND rr.rn = 1", recentReviewSub).
		Joins("LEFT JOIN (?) rcc ON rcc.storeId = stores.id", recentCertCountSub).
		Preload("Type", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, description")
		})

	switch filter {
	case "popular":
		queryBuilder = queryBuilder.
			Order("rc.review_count DESC").
			Order("stores.averageRating DESC").
			Order("rcc.recent_cert_count DESC").
			Order("stores.id ASC")
	case "recent_cert":
		queryBuilder = queryBuilder.
			Order("rcc.recent_cert_count DESC").
			Order("stores.id ASC")
	default:
		queryBuilder = queryBuilder.Order("stores.id ASC")
	}

	err := queryBuilder.Find(&stores).Error
	if err != nil {
		return dto.FindInBoundsDto{}, err
	}

	data := make([]dto.StoreInBoundsResponse, 0, len(stores))
	for _, store := range stores {
		var storeType dto.StoreTypeResponse
		if store.Type != nil {
			storeType = dto.StoreTypeResponse{
				ID:          store.Type.ID,
				Name:        store.Type.Name,
				Description: store.Type.Description,
			}
		}

		data = append(data, dto.StoreInBoundsResponse{
			ID:            store.ID,
			Name:          store.Name,
			Address:       store.Address,
			Region1:       store.Region1,
			Region2:       store.Region2,
			Latitude:      store.Latitude,
			Longitude:     store.Longitude,
			Phone:         store.Phone,
			AverageRating: store.AverageRating,
			Type:          storeType,
			CreatedAt:     store.CreatedAt,
			UpdatedAt:     store.UpdatedAt,
			RecentReview:  store.RecentReview,
			ReviewCount:   store.ReviewCount,
		})
	}

	return dto.FindInBoundsDto{
		Success: true,
		Data:    data,
		Meta: dto.Meta{
			Count: total,
		},
	}, nil
}
