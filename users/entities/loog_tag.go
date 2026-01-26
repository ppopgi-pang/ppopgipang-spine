package entities

import "time"

type LootTag struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`             // 태그명 (예: 인형, 피규어)
	IconName  *string   `gorm:"column:iconName;type:varchar(100)" json:"iconName"` // 아이콘 이미지 경로
	SortOrder int       `gorm:"column:sortOrder;default:0;index:idx_loot_tags_active" json:"sortOrder"`
	IsActive  int8      `gorm:"column:isActive;type:tinyint;default:1;index:idx_loot_tags_active" json:"isActive"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
}

func (LootTag) TableName() string {
	return "loot_tags"
}
