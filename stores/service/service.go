package service

import (
	"context"
	"errors"
	"math"

	certEntities "github.com/ppopgi-pang/ppopgipang-spine/certifications/entities"
	reviewDtos "github.com/ppopgi-pang/ppopgipang-spine/reviews/dto"
	reviewEntities "github.com/ppopgi-pang/ppopgipang-spine/reviews/entities"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	userEntities "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
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
			"JOIN "+userEntities.UserStoreBookmark{}.TableName()+" usb ON usb.storeId = stores.id AND usb.userId = ?",
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
		Table(reviewEntities.Review{}.TableName()).
		Select("storeId, COUNT(*) as review_count").
		Group("storeId")

	recentReviewSub := s.db.
		Table(reviewEntities.Review{}.TableName()).
		Select("storeId, content, ROW_NUMBER() OVER (PARTITION BY storeId ORDER BY createdAt DESC, id DESC) as rn")

	recentCertCountSub := s.db.
		Table(certEntities.Certification{}.TableName()).
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
			"JOIN "+userEntities.UserStoreBookmark{}.TableName()+" usb ON usb.storeId = stores.id AND usb.userId = ?",
			*userID,
		)
	}

	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return dto.FindInBoundsDto{}, err
	}

	reviewCountSub := s.db.
		Table(reviewEntities.Review{}.TableName()).
		Select("storeId, COUNT(*) as review_count").
		Group("storeId")

	recentReviewSub := s.db.
		Table(reviewEntities.Review{}.TableName()).
		Select("storeId, content, ROW_NUMBER() OVER (PARTITION BY storeId ORDER BY createdAt DESC, id DESC) as rn")

	recentCertCountSub := s.db.
		Table(certEntities.Certification{}.TableName()).
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

func (s *StoreService) SearchStore(ctx context.Context, keyword string, latitude, longitude float64, page, size int) (dto.StoreSearchResponse, error) {
	type storeWithDistance struct {
		entities.Store `gorm:"embedded"`
		Distance       float64 `gorm:"column:distance"`
	}

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	likeKeyword := "%" + keyword + "%"

	baseQuery := s.db.WithContext(ctx).
		Model(&entities.Store{}).
		Where(
			"(stores.name LIKE ? OR stores.address LIKE ?)",
			likeKeyword,
			likeKeyword,
		)

	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return dto.StoreSearchResponse{}, err
	}

	queryBuilder := baseQuery.
		Session(&gorm.Session{}).
		Preload("Photos", "`type` = ?", "cover").
		Preload("Type")

	if latitude != 0 && longitude != 0 {
		queryBuilder = queryBuilder.
			Select(
				"stores.*, ST_Distance_Sphere(POINT(?, ?), POINT(stores.longitude, stores.latitude)) as distance",
				longitude,
				latitude,
			).
			Order("distance ASC")
	} else {
		queryBuilder = queryBuilder.
			Select("stores.*, 0 as distance").
			Order("stores.name ASC")
	}

	var rows []storeWithDistance
	if err := queryBuilder.
		Offset((page - 1) * size).
		Limit(size).
		Find(&rows).
		Error; err != nil {
		return dto.StoreSearchResponse{}, err
	}

	data := make([]dto.StoreResponse, 0, len(rows))
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

		var thumbnailName *string
		if len(store.Photos) > 0 {
			thumbnailName = store.Photos[0].ImageName
		}

		data = append(data, dto.StoreResponse{
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
			CreatedAt:     store.CreatedAt,
			UpdatedAt:     store.UpdatedAt,
			// 첫번째 커버 사진
			ThumbnailName: thumbnailName,
		})
	}

	return dto.StoreSearchResponse{
		Success: true,
		Data:    data,
		Meta: dto.Meta{
			Count: total,
		},
	}, nil
}

func (s *StoreService) FindByStoreSummaryId(ctx context.Context, storeId int64) (dto.StoreSummaryResponse, error) {
	db := s.db.WithContext(ctx)

	var result entities.Store
	err := db.Preload("Photos").First(&result, "id = ?", storeId).Error
	if err != nil {
		return dto.StoreSummaryResponse{}, err
	}

	var reviews []reviewEntities.Review
	err = db.
		Model(&reviewEntities.Review{}).
		Where("storeId = ?", result.ID).
		Scan(&reviews).
		Error
	if err != nil {
		return dto.StoreSummaryResponse{}, err
	}

	imageNames := make([]string, 0, len(result.Photos))
	for _, photo := range result.Photos {
		imageNames = append(imageNames, *photo.ImageName)
	}
	return dto.StoreSummaryResponse{
		ID:            result.ID,
		Name:          result.Name,
		AverageRating: result.AverageRating,
		ImageNames:    imageNames,
		ReviewCount:   len(reviews),
	}, nil
}

func (s *StoreService) FindByStoreDetailId(ctx context.Context, storeId int64, userId *int64) (dto.StoreDetailResponse, error) {
	db := s.db.WithContext(ctx)

	var store entities.Store
	err := db.
		Model(&entities.Store{}).
		Preload("OpeningHours", func(db *gorm.DB) *gorm.DB {
			return db.Order("dayOfWeek ASC")
		}).
		Preload("Facilities").
		First(&store, "id = ?", storeId).
		Error
	if err != nil {
		return dto.StoreDetailResponse{}, err
	}

	isBookmark := false
	if userId != nil {
		var userStoreBookmark userEntities.UserStoreBookmark
		bookmarkErr := db.
			Model(&userEntities.UserStoreBookmark{}).
			First(&userStoreBookmark, "userId = ? AND storeId = ?", *userId, storeId).
			Error
		if bookmarkErr == nil {
			isBookmark = true
		} else if !errors.Is(bookmarkErr, gorm.ErrRecordNotFound) {
			return dto.StoreDetailResponse{}, bookmarkErr
		}
	}

	openingHours := make([]dto.StoreOpeningHourResponse, 0, len(store.OpeningHours))
	for _, item := range store.OpeningHours {
		openingHours = append(openingHours, dto.StoreOpeningHourResponse{
			ID:        item.ID,
			DayOfWeek: item.DayOfWeek,
			OpenTime:  item.OpenTime,
			CloseTime: item.CloseTime,
			IsClosed:  item.IsClosed,
		})
	}

	facilityResponse := dto.StoreFacilityResponse{}
	if store.Facilities != nil {
		facilityResponse = dto.StoreFacilityResponse{
			MachineCount:   store.Facilities.MachineCount,
			PaymentMethods: []string(store.Facilities.PaymentMethods),
		}
	}

	return dto.StoreDetailResponse{
		IsBookmark:                isBookmark,
		StoreOpeningHourResponses: openingHours,
		Phone:                     store.Phone,
		StoreFacilityResponse:     facilityResponse,
	}, nil
}

func (s *StoreService) GetStoreStatById(ctx context.Context, storeId int64, userId *int64) (dto.VisitHistoryResponse, error) {
	db := s.db.WithContext(ctx)

	// UserID, StoreID로 MyStat 조회
	var myStat *dto.MyStat
	if userId != nil {
		var myStoreStat userEntities.UserStoreStat
		err := db.
			Model(&userEntities.UserStoreStat{}).
			First(&myStoreStat, "userId = ? AND storeId = ?", *userId, storeId).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.VisitHistoryResponse{}, err
		}
		if err == nil {
			myStat = &dto.MyStat{
				VisitCount: myStoreStat.VisitCount,
			}
		}
	}
	// StoreID로 Stat 집계하여 평균 방문수, 최대 방문수, 한달 방문수 추출
	type storeStatAggregate struct {
		AverageVisitCount   float64 `gorm:"column:averageVisitCount"`
		MaxVisitCount       int     `gorm:"column:maxVisitCount"`
		MonthlyVisitorCount int     `gorm:"column:monthlyVisitorCount"`
	}
	var aggregate storeStatAggregate
	err := db.
		Model(&userEntities.UserStoreStat{}).
		Select(
			"COALESCE(AVG(visitCount), 0) as averageVisitCount, "+
				"COALESCE(MAX(visitCount), 0) as maxVisitCount, "+
				"COALESCE(SUM(CASE WHEN lastVisitedAt >= DATE_SUB(NOW(), INTERVAL 1 MONTH) THEN 1 ELSE 0 END), 0) as monthlyVisitorCount",
		).
		Where("storeId = ?", storeId).
		Scan(&aggregate).Error
	if err != nil {
		return dto.VisitHistoryResponse{}, err
	}

	// StoreID로 Review 조회
	var reviews []reviewEntities.Review
	err = db.
		Model(&reviewEntities.Review{}).
		Where("storeId = ?", storeId).
		Order("createdAt DESC, id DESC").
		Find(&reviews).Error
	if err != nil {
		return dto.VisitHistoryResponse{}, err
	}

	reviewImages := make([]string, 0)
	reviewResponses := make([]reviewDtos.ReviewResponse, 0, len(reviews))
	for _, review := range reviews {
		reviewResponses = append(reviewResponses, reviewDtos.ReviewResponse{
			ID:        review.ID,
			Rating:    review.Rating,
			Content:   review.Content,
			Images:    []string(review.Images),
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
		reviewImages = append(reviewImages, []string(review.Images)...)
	}

	return dto.VisitHistoryResponse{
		MyStat: myStat,
		OtherUserStat: dto.OtherUserStat{
			AverageVisitCount:   int(math.Round(aggregate.AverageVisitCount)),
			MaxVisitCount:       aggregate.MaxVisitCount,
			MonthlyVisitorCount: aggregate.MonthlyVisitorCount,
		},
		ReviewImages:    reviewImages,
		ReviewResponses: reviewResponses,
	}, nil
}

func (s *StoreService) GetStoreReviewsById(ctx context.Context, storeId int64) (dto.StoreReviewResponse, error) {
	db := s.db.WithContext(ctx)

	var reviews []reviewEntities.Review
	err := db.
		Model(&reviewEntities.Review{}).
		Where("storeId = ?", storeId).
		Order("createdAt DESC, id DESC").
		Find(&reviews).Error
	if err != nil {
		return dto.StoreReviewResponse{}, err
	}

	reviewImages := make([]string, 0)
	reviewResponses := make([]reviewDtos.ReviewResponse, 0, len(reviews))
	for _, review := range reviews {
		reviewResponses = append(reviewResponses, reviewDtos.ReviewResponse{
			ID:        review.ID,
			Rating:    review.Rating,
			Content:   review.Content,
			Images:    []string(review.Images),
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
		reviewImages = append(reviewImages, []string(review.Images)...)
	}

	return dto.StoreReviewResponse{
		ReviewImages:    reviewImages,
		ReviewResponses: reviewResponses,
	}, nil
}
