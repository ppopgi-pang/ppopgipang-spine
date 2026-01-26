package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type TradeChatMessage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomID    uint      `gorm:"column:roomId;not null" json:"roomId"`
	SenderID  uint      `gorm:"column:senderId;not null" json:"senderId"`
	Message   *string   `gorm:"type:text" json:"message"`
	ImageName *string   `gorm:"column:imageName;type:varchar(255)" json:"imageName"`
	IsRead    int8      `gorm:"column:isRead;type:tinyint;default:0" json:"isRead"`
	SentAt    time.Time `gorm:"column:sentAt;type:datetime(6);autoCreateTime" json:"sentAt"`

	// Associations
	Room   *TradeChatRoom `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE" json:"room,omitempty"`
	Sender *users.User    `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE" json:"sender,omitempty"`
}

func (TradeChatMessage) TableName() string {
	return "trade_chat_message"
}
