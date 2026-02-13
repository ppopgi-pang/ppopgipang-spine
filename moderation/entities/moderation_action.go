package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type ModerationAction struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ReportID  *int64    `gorm:"column:reportId" json:"reportId"`
	AdminID   *int64    `gorm:"column:adminId" json:"adminId"`
	Action    *string   `gorm:"type:varchar(50)" json:"action"` // ban_user, delete_content, etc
	Note      *string   `gorm:"type:text" json:"note"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	Report *ContentReport `gorm:"foreignKey:ReportID;constraint:OnDelete:CASCADE" json:"report,omitempty"`
	Admin  *users.User    `gorm:"foreignKey:AdminID;constraint:OnDelete:CASCADE" json:"admin,omitempty"`
}

func (ModerationAction) TableName() string {
	return "moderation_actions"
}
