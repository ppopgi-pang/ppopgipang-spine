package entities

import (
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/commons/types"
	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type Notification struct {
	ID        int64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    *int64        `gorm:"column:userId;index:idx_noti_user_read" json:"userId"`
	Type      *string       `gorm:"type:varchar(50)" json:"type"` // level_up, trade_msg, store_hot, etc
	Title     *string       `gorm:"type:varchar(100)" json:"title"`
	Message   *string       `gorm:"type:varchar(255)" json:"message"`
	Payload   types.JSONMap `gorm:"type:json" json:"payload"` // 클릭 시 이동할 링크 정보 등
	IsRead    int8          `gorm:"column:isRead;type:tinyint;default:0;index:idx_noti_user_read" json:"isRead"`
	CreatedAt time.Time     `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	User *users.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}
