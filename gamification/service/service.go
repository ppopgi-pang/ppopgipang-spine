package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	certEntity "github.com/ppopgi-pang/ppopgipang-spine/certifications/entities"
	"github.com/ppopgi-pang/ppopgipang-spine/commons/types"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/gamification/entities"
	reviewEntity "github.com/ppopgi-pang/ppopgipang-spine/reviews/entities"
	userEntity "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
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
		var progress userEntity.UserProgress
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
			Model(&certEntity.Certification{}).
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
			Model(&reviewEntity.Review{}).
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
