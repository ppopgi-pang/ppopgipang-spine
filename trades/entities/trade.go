package entities

import (
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/commons/types"
	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type Trade struct {
	ID          uint                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       *string               `gorm:"type:varchar(100)" json:"title"`
	Description *string               `gorm:"type:text" json:"description"`
	Images      types.JSONStringArray `gorm:"type:json" json:"images"` // 이미지 URL 배열
	Price       *int                  `gorm:"type:int" json:"price"`
	Type        string                `gorm:"type:enum('sale','exchange');not null" json:"type"`
	Status      string                `gorm:"type:enum('active','reserved','completed','cancelled');default:active" json:"status"`
	CreatedAt   time.Time             `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time             `gorm:"column:updatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`
	UserID      uint                  `gorm:"column:userId;not null" json:"userId"` // 판매자
	LootID      *uint                 `gorm:"column:lootId" json:"lootId"`          // user_loots 테이블의 아이템을 파는 경우 연결

	// Associations
	User      *users.User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Loot      *users.UserLoot `gorm:"foreignKey:LootID;constraint:OnDelete:SET NULL" json:"loot,omitempty"`
	ChatRooms []TradeChatRoom `gorm:"foreignKey:TradeID" json:"chatRooms,omitempty"`
}

func (Trade) TableName() string {
	return "trades"
}
