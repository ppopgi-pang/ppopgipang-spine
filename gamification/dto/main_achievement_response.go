package dto

import "time"

type AchievementProgressResponse struct {
	AchievementID   int64      `json:"achievementId" example:"101"`
	Code            *string    `json:"code" example:"LOGIN_STREAK"`
	Name            *string    `json:"name" example:"연속 접속 7일"`
	Description     *string    `json:"description" example:"7일 연속으로 접속하세요"`
	BadgeImageName  *string    `json:"badgeImageName" example:"badge_login_7.png"`
	IsCompleted     bool       `json:"isCompleted" example:"false"`
	EarnedAt        *time.Time `json:"earnedAt,omitempty" example:"2024-01-10T08:15:00Z"`
	ProgressCurrent int        `json:"progressCurrent" example:"3"`
	ProgressTarget  int        `json:"progressTarget" example:"7"`
	Remaining       int        `json:"remaining" example:"4"`
	ActionLabel     string     `json:"actionLabel" example:"4일 더 접속"`
}

type GamificationMainAchievementResponse struct {
	Success bool                         `json:"success" example:"true"`
	Item    *AchievementProgressResponse `json:"item"`
}
