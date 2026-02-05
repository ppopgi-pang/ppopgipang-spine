package entities

import (
	stores "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
)

type Stamp struct {
	ID        int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID   *int64  `gorm:"column:storeId" json:"storeId"`
	ImageName *string `gorm:"column:imageName;type:varchar(255)" json:"imageName"` // 가게 로고 혹은 지역 심볼

	// Associations
	Store      *stores.Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
	UserStamps []UserStamp   `gorm:"foreignKey:StampID" json:"userStamps,omitempty"`
}

func (Stamp) TableName() string {
	return "stamps"
}
