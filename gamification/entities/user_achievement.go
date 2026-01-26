package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type UserAchievement struct {
	UserID        uint       `gorm:"column:userId;primaryKey" json:"userId"`
	AchievementID uint       `gorm:"column:achievementId;primaryKey" json:"achievementId"`
	EarnedAt      *time.Time `gorm:"column:earnedAt;type:datetime(6)" json:"earnedAt"`

	// Associations
	User        *users.User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Achievement *Achievement `gorm:"foreignKey:AchievementID;constraint:OnDelete:CASCADE" json:"achievement,omitempty"`
}

func (UserAchievement) TableName() string {
	return "user_achievements"
}
