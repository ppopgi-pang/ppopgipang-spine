package entities

import (
	"time"

	stores "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type UserStoreStat struct {
	UserID        uint       `gorm:"column:userId;primaryKey" json:"userId"`
	StoreID       uint       `gorm:"column:storeId;primaryKey" json:"storeId"`
	VisitCount    int        `gorm:"column:visitCount;default:0" json:"visitCount"`
	LootCount     int        `gorm:"column:lootCount;default:0" json:"lootCount"`
	LastVisitedAt *time.Time `gorm:"column:lastVisitedAt;type:datetime(6)" json:"lastVisitedAt"`
	IsScrapped    int8       `gorm:"column:isScrapped;type:tinyint(1);not null;default:0" json:"isScrapped"`
	Tier          string     `gorm:"type:enum('unknown','visited','master');default:unknown" json:"tier"`

	// Associations
	User  *users.User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Store *stores.Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
}

func (UserStoreStat) TableName() string {
	return "user_store_stats"
}
