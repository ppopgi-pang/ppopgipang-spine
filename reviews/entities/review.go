package entities

import (
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/commons/types"
	stores "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type Review struct {
	ID        uint                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Rating    int                   `gorm:"not null" json:"rating"`   // 별점
	Content   *string               `gorm:"type:text" json:"content"` // 리뷰 내용
	Images    types.JSONStringArray `gorm:"type:json" json:"images"`  // 리뷰 이미지
	CreatedAt time.Time             `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time             `gorm:"column:updatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`
	UserID    *uint                 `gorm:"column:userId" json:"userId"`   // 작성자 ID
	StoreID   *uint                 `gorm:"column:storeId" json:"storeId"` // 가게 ID

	// Associations
	User  *users.User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Store *stores.Store `gorm:"foreignKey:StoreID" json:"store,omitempty"`
}

func (Review) TableName() string {
	return "reviews"
}
