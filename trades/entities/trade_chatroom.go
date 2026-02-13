package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type TradeChatRoom struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TradeID   int64     `gorm:"column:tradeId;not null" json:"tradeId"`
	SellerID  int64     `gorm:"column:sellerId;not null" json:"sellerId"`
	BuyerID   int64     `gorm:"column:buyerId;not null" json:"buyerId"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`

	// Associations
	Trade    *Trade             `gorm:"foreignKey:TradeID;constraint:OnDelete:CASCADE" json:"trade,omitempty"`
	Seller   *users.User        `gorm:"foreignKey:SellerID;constraint:OnDelete:CASCADE" json:"seller,omitempty"`
	Buyer    *users.User        `gorm:"foreignKey:BuyerID;constraint:OnDelete:CASCADE" json:"buyer,omitempty"`
	Messages []TradeChatMessage `gorm:"foreignKey:RoomID" json:"messages,omitempty"`
}

func (TradeChatRoom) TableName() string {
	return "trade_chat_room"
}
