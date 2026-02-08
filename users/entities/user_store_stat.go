package entities

import (
	"time"

	stores "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
)

type UserStoreStat struct {
	UserID        int64      `gorm:"column:userId;primaryKey" json:"userId"`
	StoreID       int64      `gorm:"column:storeId;primaryKey" json:"storeId"`
	VisitCount    int        `gorm:"column:visitCount;default:0" json:"visitCount"`
	LootCount     int        `gorm:"column:lootCount;default:0" json:"lootCount"`
	LastVisitedAt *time.Time `gorm:"column:lastVisitedAt;type:datetime(6)" json:"lastVisitedAt"`
	IsScrapped    bool       `gorm:"column:isScrapped;type:boolean;not null;default:0" json:"isScrapped"`
	Tier          string     `gorm:"type:enum('unknown','visited','master');default:unknown" json:"tier"`

	// Associations
	User  *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Store *stores.Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
}

func (UserStoreStat) TableName() string {
	return "user_store_stats"
}
