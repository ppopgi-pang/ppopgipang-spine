package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type PushSubscription struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    *int64    `gorm:"column:userId" json:"userId"`
	Platform  string    `gorm:"type:enum('web','android','ios')" json:"platform"`
	Endpoint  string    `gorm:"type:varchar(500);not null" json:"endpoint"`
	AuthKey   *string   `gorm:"column:authKey;type:varchar(255)" json:"authKey"`
	P256dhKey *string   `gorm:"column:p256dhKey;type:varchar(255)" json:"p256dhKey"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	User *users.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (PushSubscription) TableName() string {
	return "push_subscriptions"
}
