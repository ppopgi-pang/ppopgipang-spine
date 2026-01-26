package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type ContentReport struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ReporterID  *uint     `gorm:"column:reporterId" json:"reporterId"`
	TargetType  string    `gorm:"column:targetType;type:enum('certification','trade','chat','store');not null" json:"targetType"`
	TargetID    uint      `gorm:"column:targetId;not null" json:"targetId"`
	Reason      *string   `gorm:"type:varchar(100)" json:"reason"`
	Description *string   `gorm:"type:text" json:"description"`
	Status      string    `gorm:"type:enum('open','resolved','rejected');default:open" json:"status"`
	CreatedAt   time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	Reporter          *users.User        `gorm:"foreignKey:ReporterID;constraint:OnDelete:CASCADE" json:"reporter,omitempty"`
	ModerationActions []ModerationAction `gorm:"foreignKey:ReportID" json:"moderationActions,omitempty"`
}

func (ContentReport) TableName() string {
	return "content_reports"
}
