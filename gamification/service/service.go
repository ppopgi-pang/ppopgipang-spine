package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	certEntities "github.com/ppopgi-pang/ppopgipang-spine/certifications/entities"
	"github.com/ppopgi-pang/ppopgipang-spine/commons/types"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/entities"
	gameEntities "github.com/ppopgi-pang/ppopgipang-spine/gamification/entities"
	reviewEntities "github.com/ppopgi-pang/ppopgipang-spine/reviews/entities"
	userEntities "github.com/ppopgi-pang/ppopgipang-spine/users/entities"

	"gorm.io/gorm"
)

type GamificationService struct {
	db *gorm.DB
}

func NewGamificationService(db *gorm.DB) *GamificationService {
	return &GamificationService{db: db}
}

func (g *GamificationService) GetMainAchievementSummary(ctx context.Context, userID int64) (dto.GamificationMainAchievementResponse, error) {
	db := g.db.WithContext(ctx)

	var achievements []entities.Achievement
	if err := db.
		Where("isHidden = ?", false).
		Order("id ASC").
		Find(&achievements).
		Error; err != nil {
		return dto.GamificationMainAchievementResponse{}, err
	}

	var userAchievements []entities.UserAchievement
	if err := db.
		Where("userId = ?", userID).
		Find(&userAchievements).
		Error; err != nil {
		return dto.GamificationMainAchievementResponse{}, err
	}

	completed := make(map[int64]struct{}, len(userAchievements))
	for _, ua := range userAchievements {
		completed[ua.AchievementID] = struct{}{}
	}

	var currentAchievement *entities.Achievement
	for i := range achievements {
		if _, ok := completed[achievements[i].ID]; ok {
			continue
		}
		currentAchievement = &achievements[i]
		break
	}

	if currentAchievement == nil {
		return dto.GamificationMainAchievementResponse{Success: true, Item: nil}, nil
	}

	stats, err := g.loadUserAchievementStats(ctx, userID, []entities.Achievement{*currentAchievement})
	if err != nil {
		return dto.GamificationMainAchievementResponse{}, err
	}

	achievementType := getAchievementType(currentAchievement.ConditionJSON)
	target := getTargetCount(currentAchievement.ConditionJSON)
	current := getCurrentProgress(stats, achievementType)

	if target > 0 && current > target {
		current = target
	}

	remaining := max(target-current, 0)

	item := dto.AchievementProgressResponse{
		AchievementID:   currentAchievement.ID,
		Code:            currentAchievement.Code,
		Name:            currentAchievement.Name,
		Description:     currentAchievement.Description,
		BadgeImageName:  currentAchievement.BadgeImageName,
		IsCompleted:     false,
		EarnedAt:        nil,
		ProgressCurrent: current,
		ProgressTarget:  target,
		Remaining:       remaining,
		ActionLabel:     buildActionLabel(currentAchievement.ConditionJSON, achievementType, remaining),
	}

	return dto.GamificationMainAchievementResponse{
		Success: true,
		Item:    &item,
	}, nil
}

func (g *GamificationService) ProcessCertification(ctx context.Context, userId, storeId int64, certificationType string, expAmount int) (dto.RewardResponse, error) {
	db := g.db.WithContext(ctx)

	var userProgress userEntities.UserProgress
	err := db.
		Where("userId = ?", userId).
		First(&userProgress).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userProgress = userEntities.UserProgress{
				UserID:     userId,
				Exp:        0,
				Level:      1,
				StreakDays: 0,
			}
		} else {
			return dto.RewardResponse{}, err
		}
	}

	oldLevel := userProgress.Level
	oldExp := userProgress.Exp
	newExp := oldExp + expAmount
	newLevel := g.calculateLevel(newExp)
	levelUp := newLevel > oldLevel

	nowTime := time.Now()
	userProgress.Exp = newExp
	userProgress.Level = newLevel
	userProgress.LastActivityAt = &nowTime

	if err := db.Save(&userProgress).Error; err != nil {
		return dto.RewardResponse{}, fmt.Errorf("UserProgress 생성 중 오류 발생")
	}

	err = g.updateStoreStats(db, userId, storeId, certificationType)
	if err != nil {
		return dto.RewardResponse{}, fmt.Errorf("유저 가게 상태 업데이트 실패")
	}

	newStamp, err := g.checkAndGrantStamp(db, userId, storeId)
	if err != nil {
		return dto.RewardResponse{}, fmt.Errorf("스탬프 지급 체크 실패")
	}
	newBadges, err := g.checkAndGrantBadge(db, userId, certificationType)
	if err != nil {
		return dto.RewardResponse{}, fmt.Errorf("배지 달성 지급 실패")
	}
	err = g.updateStreak(db, userProgress)
	if err != nil {
		return dto.RewardResponse{}, fmt.Errorf("연속 방문일 업데이트 실패")
	}
	expToNextLevel := (newLevel * 100) - newExp

	var resultNewLevel *int
	if levelUp == true {
		resultNewLevel = &newLevel
	}

	return dto.RewardResponse{
		Exp:            expAmount,
		TotalExp:       newExp,
		CurrentLevel:   newLevel,
		LevelUp:        levelUp,
		ExpToNextLevel: expToNextLevel,
		NewLevel:       resultNewLevel,
		NewStamp:       newStamp,
		NewBadges:      &newBadges,
	}, nil

}

// 레벨 계산
func (g *GamificationService) calculateLevel(totalExp int) int {
	return totalExp/100 + 1
}

// 유저의 가게 상태 업데이트
func (g *GamificationService) updateStoreStats(db *gorm.DB, userId int64, storeId int64, certificationType string) error {
	var stat userEntities.UserStoreStat
	err := db.
		Where("userId = ? AND storeId = ?", userId, storeId).
		First(&stat).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var lootCount int
			if certificationType == "loot" {
				lootCount = 1
			} else {
				lootCount = 0
			}

			nowTime := time.Now()

			stat = userEntities.UserStoreStat{
				UserID:        userId,
				StoreID:       storeId,
				VisitCount:    1,
				LootCount:     lootCount,
				LastVisitedAt: &nowTime,
				Tier:          "visited",
				IsScrapped:    false,
			}
		} else {
			return err
		}
	} else {
		stat.VisitCount += 1
		if certificationType == "loot" {
			stat.LootCount += 1
		}
		nowTime := time.Now()
		stat.LastVisitedAt = &nowTime

		// 티어 업데이트
		if stat.VisitCount >= 5 {
			stat.Tier = "master"
		} else {
			stat.Tier = "visited"
		}
	}

	if err := db.Save(&stat).Error; err != nil {
		return fmt.Errorf("UserStatUser 생성 실패")
	}
	return nil
}

// 스탬프 지급 체크
func (g *GamificationService) checkAndGrantStamp(db *gorm.DB, userId, storedId int64) (*dto.StampResponse, error) {
	// 해당 가게의 이전 인증 횟수 확인
	var previousCertCount int64
	err := db.
		Model(&certEntities.Certification{}).
		Where("userId = ? AND storeId = ?", userId, storedId).
		Count(&previousCertCount).Error
	if err != nil {
		return nil, err
	}

	// 첫 방문이 아니면 스탬프 지급을 하지 않음
	if previousCertCount > 0 {
		return nil, nil
	}

	// 해당 가게의 스탬프 찾기
	var stamp gameEntities.Stamp
	err = db.
		Where("storeId = ?", storedId).
		Preload("Store").
		First(&stamp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	// 이미 획득한 스탬프인지 확인
	var existingUserStamp gameEntities.UserStamp
	err = db.
		Where("userId = ? AND stampId = ?", userId, stamp.ID).
		First(&existingUserStamp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// not acquired yet -> continue to grant
		} else {
			return nil, err
		}
	} else {
		// already acquired
		return nil, nil
	}

	// 스탬프 지급
	nowTime := time.Now()
	userStamp := gameEntities.UserStamp{
		UserID:     userId,
		StampID:    stamp.ID,
		AcquiredAt: &nowTime,
	}

	if err := db.Create(&userStamp).Error; err != nil {
		return nil, fmt.Errorf("UserStamp 지급 실패")
	}
	var storeName *string
	if stamp.Store != nil {
		storeName = &stamp.Store.Name
	}
	return &dto.StampResponse{
		ID:        stamp.ID,
		ImageName: stamp.ImageName,
		StoreName: storeName,
	}, nil

}

// 연속 방문일 업데이트
func (g *GamificationService) updateStreak(db *gorm.DB, userProgress userEntities.UserProgress) error {
	nowTime := time.Now()
	lastActivity := userProgress.LastActivityAt

	if lastActivity == nil {
		userProgress.StreakDays = 1
		return nil
	}

	daysDiff := int(
		nowTime.Sub(*lastActivity).Hours() / 24,
	)

	switch daysDiff {
	case 0:
		// 같은 날은 유지
		return nil
	case 1:
		// 연속 방문
		userProgress.StreakDays += 1
	default:
		userProgress.StreakDays = 1
	}

	err := db.Save(&userProgress).Error
	if err != nil {
		return err
	}
	return nil
}

// 배지 달성 체크 & 지급
func (g *GamificationService) checkAndGrantBadge(db *gorm.DB, userId int64, certificationType string) ([]dto.NewBadgeResponse, error) {
	var newBadges []dto.NewBadgeResponse
	var userAchievements []gameEntities.UserAchievement
	err := db.
		Where("userId = ?", userId).
		Find(&userAchievements).Error
	if err != nil {
		return []dto.NewBadgeResponse{}, err
	}

	seen := make(map[int64]struct{}, len(userAchievements))
	for _, achievement := range userAchievements {
		id := achievement.AchievementID
		if _, exists := seen[id]; exists {
			continue
		}
		seen[id] = struct{}{}
	}

	badgesToCheck := []string{
		"FIRST_LOOT",
		"FIRST_CHECKIN",
		"LOOT_10",
		"VISIT_5_STORES",
		"STREAK_7",
	}

	for _, badgeCode := range badgesToCheck {
		var achievement gameEntities.Achievement
		err := db.
			Where("code = ?", badgeCode).
			First(&achievement).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			continue
		}
		if err != nil {
			return []dto.NewBadgeResponse{}, err
		}

		if _, exists := seen[achievement.ID]; exists {
			continue
		}

		var shouldGrant bool

		switch badgeCode {
		case "FIRST_LOOT":
			if certificationType == "loot" {
				var lootCount int64
				err := db.Model(&certEntities.Certification{}).
					Where("userId = ? AND type = 'loot'", userId).
					Count(&lootCount).Error
				if err != nil {
					return []dto.NewBadgeResponse{}, err
				}
				shouldGrant = lootCount == 1
			}
		case "FIRST_CHECKIN":
			if certificationType == "checkin" {
				var checkinCount int64
				err := db.Model(&certEntities.Certification{}).
					Where("userId = ? AND type = 'checkin'", userId).
					Count(&checkinCount).Error
				if err != nil {
					return []dto.NewBadgeResponse{}, err
				}
				shouldGrant = checkinCount == 1
			}
		case "LOOT_10":
			var lootCount int64
			err := db.Model(&certEntities.Certification{}).
				Where("userId = ? AND type = 'loot'", userId).
				Count(&lootCount).Error
			if err != nil {
				return []dto.NewBadgeResponse{}, err
			}
			shouldGrant = lootCount >= 10
		case "VISIT_5_STORES":
			var visitedStoresCount int64
			err := db.Model(&userEntities.UserStoreStat{}).
				Where("userId = ? AND visitCount > 0", userId).
				Count(&visitedStoresCount).Error
			if err != nil {
				return []dto.NewBadgeResponse{}, err
			}
			shouldGrant = visitedStoresCount >= 5
		case "STREAK_7":
			var userProgress userEntities.UserProgress
			err := db.
				Where("userId = ?", userId).
				First(&userProgress).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				shouldGrant = false
				break
			}
			if err != nil {
				return []dto.NewBadgeResponse{}, err
			}
			shouldGrant = userProgress.StreakDays >= 7
		default:
			return []dto.NewBadgeResponse{}, fmt.Errorf("예외 케이스 발생")
		}

		if shouldGrant == true {
			nowTime := time.Now()
			userAchievement := gameEntities.UserAchievement{
				UserID:        userId,
				AchievementID: achievement.ID,
				EarnedAt:      &nowTime,
			}
			if err := db.Create(&userAchievement).Error; err != nil {
				return []dto.NewBadgeResponse{}, err
			}

			newBadges = append(newBadges, dto.NewBadgeResponse{
				ID:             achievement.ID,
				Code:           achievement.Code,
				Name:           achievement.Name,
				Description:    achievement.Description,
				BadgeImageName: achievement.BadgeImageName,
			})
		}
	}

	return newBadges, nil
}

type achievementStats struct {
	streakDays         int
	stampCount         int
	certificationCount int
	reviewCount        int
}

func (g *GamificationService) loadUserAchievementStats(ctx context.Context, userID int64, achievements []entities.Achievement) (achievementStats, error) {
	needed := make(map[string]struct{}, len(achievements))
	for _, achievement := range achievements {
		if achievementType := getAchievementType(achievement.ConditionJSON); achievementType != "" {
			needed[achievementType] = struct{}{}
		}
	}

	stats := achievementStats{}
	db := g.db.WithContext(ctx)

	if _, ok := needed["STREAK_DAYS"]; ok {
		var progress userEntities.UserProgress
		err := db.First(&progress, "userId = ?", userID).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return achievementStats{}, err
		}
		if err == nil {
			stats.streakDays = progress.StreakDays
		}
	}

	if _, ok := needed["STAMP_COUNT"]; ok {
		var count int64
		if err := db.
			Model(&entities.UserStamp{}).
			Where("userId = ?", userID).
			Count(&count).
			Error; err != nil {
			return achievementStats{}, err
		}
		stats.stampCount = int(count)
	}

	if _, ok := needed["CERTIFICATION_COUNT"]; ok {
		var count int64
		if err := db.
			Model(&certEntities.Certification{}).
			Where("userId = ?", userID).
			Count(&count).
			Error; err != nil {
			return achievementStats{}, err
		}
		stats.certificationCount = int(count)
	}

	if _, ok := needed["REVIEW_COUNT"]; ok {
		var count int64
		if err := db.
			Model(&reviewEntities.Review{}).
			Where("userId = ?", userID).
			Count(&count).
			Error; err != nil {
			return achievementStats{}, err
		}
		stats.reviewCount = int(count)
	}

	return stats, nil
}

func getAchievementType(condition types.JSONMap) string {
	if condition == nil {
		return ""
	}
	if value, ok := condition["type"]; ok {
		if str, ok := value.(string); ok {
			return strings.ToUpper(str)
		}
	}
	return ""
}

func getTargetCount(condition types.JSONMap) int {
	if condition == nil {
		return 0
	}
	if count, ok := getIntFromKeys(condition, "target", "targetCount", "count", "goal"); ok {
		return count
	}
	return 0
}

func getCurrentProgress(stats achievementStats, achievementType string) int {
	switch achievementType {
	case "STREAK_DAYS":
		return stats.streakDays
	case "STAMP_COUNT":
		return stats.stampCount
	case "CERTIFICATION_COUNT":
		return stats.certificationCount
	case "REVIEW_COUNT":
		return stats.reviewCount
	default:
		return 0
	}
}

func buildActionLabel(condition types.JSONMap, achievementType string, remaining int) string {
	if condition != nil {
		if raw, ok := getStringFromKeys(condition, "actionLabel", "action", "label"); ok {
			return formatActionLabel(raw, remaining)
		}
	}

	switch achievementType {
	case "STREAK_DAYS":
		return fmt.Sprintf("%d일 더 접속", remaining)
	case "STAMP_COUNT":
		return fmt.Sprintf("%d개 더 모으기", remaining)
	case "CERTIFICATION_COUNT":
		return fmt.Sprintf("%d회 더 인증", remaining)
	case "REVIEW_COUNT":
		return fmt.Sprintf("%d개 더 작성", remaining)
	default:
		return fmt.Sprintf("%d 더 진행", remaining)
	}
}

func formatActionLabel(label string, remaining int) string {
	if strings.Contains(label, "{remaining}") {
		return strings.ReplaceAll(label, "{remaining}", fmt.Sprintf("%d", remaining))
	}
	if strings.Contains(label, "%d") {
		return fmt.Sprintf(label, remaining)
	}
	if label == "" {
		return label
	}
	return fmt.Sprintf("%d%s", remaining, label)
}

func getIntFromKeys(condition types.JSONMap, keys ...string) (int, bool) {
	for _, key := range keys {
		if value, ok := condition[key]; ok {
			switch v := value.(type) {
			case int:
				return v, true
			case int64:
				return int(v), true
			case float64:
				return int(v), true
			case float32:
				return int(v), true
			case string:
				var parsed int
				if _, err := fmt.Sscanf(v, "%d", &parsed); err == nil {
					return parsed, true
				}
			}
		}
	}
	return 0, false
}

func getStringFromKeys(condition types.JSONMap, keys ...string) (string, bool) {
	for _, key := range keys {
		if value, ok := condition[key]; ok {
			if str, ok := value.(string); ok && str != "" {
				return str, true
			}
		}
	}
	return "", false
}
