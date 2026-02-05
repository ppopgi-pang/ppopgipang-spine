package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type UserStamp struct {
	UserID     int64      `gorm:"column:userId;primaryKey" json:"userId"`
	StampID    int64      `gorm:"column:stampId;primaryKey" json:"stampId"`
	AcquiredAt *time.Time `gorm:"column:acquiredAt;type:datetime(6)" json:"acquiredAt"`

	// Associations
	User  *users.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Stamp *Stamp      `gorm:"foreignKey:StampID;constraint:OnDelete:CASCADE" json:"stamp,omitempty"`
}

func (UserStamp) TableName() string {
	return "user_stamps"
}
