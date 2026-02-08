package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/certifications/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/certifications/entities"
	gameService "github.com/ppopgi-pang/ppopgipang-spine/gamification/service"
	storeEntities "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	userEntities "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
	"gorm.io/gorm"
)

type CertificationService struct {
	db                  *gorm.DB
	gamificationService *gameService.GamificationService
}

func NewCertificationService(db *gorm.DB, gamificationService *gameService.GamificationService) *CertificationService {
	return &CertificationService{
		db:                  db,
		gamificationService: gamificationService,
	}
}

func (s *CertificationService) CreateCheckInCertification(ctx context.Context, userId int64, input *dto.CreateCheckInRequest) (dto.CertificationResponse, error) {
	db := s.db.WithContext(ctx)

	if userId <= 0 {
		return dto.CertificationResponse{}, fmt.Errorf("유효하지 않은 userId 입니다.")
	}
	if input.StoreId <= 0 {
		return dto.CertificationResponse{}, fmt.Errorf("유효하지 않은 storeId 입니다.")
	}

	// User 존재 확인
	var userCount int64
	err := db.
		Model(&userEntities.User{}).
		Where("id = ?", userId).
		Count(&userCount).Error
	if err != nil {
		return dto.CertificationResponse{}, fmt.Errorf("user 조회 중 오류 발생: %w", err)
	}
	if userCount == 0 {
		return dto.CertificationResponse{}, fmt.Errorf("User가 존재하지 않습니다.")
	}

	// Store 존재 확인
	var storeCount int64
	err = db.
		Model(&storeEntities.Store{}).
		Where("id = ?", input.StoreId).
		Count(&storeCount).Error
	if err != nil {
		return dto.CertificationResponse{}, fmt.Errorf("store 조회 중 오류 발생: %w", err)
	}
	if storeCount == 0 {
		return dto.CertificationResponse{}, fmt.Errorf("Store가 존재하지 않습니다.")
	}

	// Certification 레코드 생성
	certification := entities.Certification{
		UserID:     userId,
		StoreID:    input.StoreId,
		Type:       "checkin",
		OccurredAt: time.Now(),
		Latitude:   input.Latitude,
		Longitude:  input.Longitude,
		Exp:        10,
		Rating:     input.Rating,
	}
	if err := db.Create(&certification).Error; err != nil {
		return dto.CertificationResponse{}, fmt.Errorf("certification 저장 중 오류 발생: %w", err)
	}

	// 별로였던 이유 연결
	if len(input.ReasonIds) > 0 {
		var presets []entities.CheckinReasonPreset
		if err := db.Where("id IN ?", input.ReasonIds).Find(&presets).Error; err != nil {
			return dto.CertificationResponse{}, fmt.Errorf("방문이 별로였던 이유를 불러오지 못했습니다: %w", err)
		}

		if err := db.Model(&certification).Association("Reasons").Replace(presets); err != nil {
			return dto.CertificationResponse{}, fmt.Errorf("방문 이유 연결 중 오류 발생: %w", err)
		}
	}

	rewards, err := s.gamificationService.ProcessCertification(ctx, userId, input.StoreId, "checkin", 10)
	if err != nil {
		return dto.CertificationResponse{}, fmt.Errorf("게이미피케이션 처리가 실패했습니다 %w", err)
	}

	return dto.CertificationResponse{
		ID:      certification.ID,
		Type:    "checkin",
		Rewards: rewards,
	}, nil
}
