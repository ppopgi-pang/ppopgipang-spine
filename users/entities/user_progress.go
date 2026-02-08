package entities

import "time"

type UserProgress struct {
	UserID         int64      `gorm:"column:userId;primaryKey" json:"userId"`
	Level          int        `gorm:"default:1" json:"level"`
	Exp            int        `gorm:"default:0" json:"exp"`
	StreakDays     int        `gorm:"column:streakDays;default:0" json:"streakDays"`
	LastActivityAt *time.Time `gorm:"column:lastActivityAt;type:datetime(6)" json:"lastActivityAt"`

	// Associations
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (UserProgress) TableName() string {
	return "user_progress"
}
