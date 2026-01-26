package entities

import "time"

type CheckinReasonPreset struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"type:varchar(100);not null" json:"content"` // 프리셋 내용
	SortOrder int       `gorm:"column:sortOrder;default:0;index:idx_checkin_reason_presets_active" json:"sortOrder"`
	IsActive  int8      `gorm:"column:isActive;type:tinyint;default:1;index:idx_checkin_reason_presets_active" json:"isActive"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	Certifications []Certification `gorm:"many2many:certification_reasons" json:"certifications,omitempty"`
}

func (CheckinReasonPreset) TableName() string {
	return "checkin_reason_presets"
}
