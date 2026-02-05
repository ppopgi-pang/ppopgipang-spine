package entities

import "github.com/ppopgi-pang/ppopgipang-spine/commons/types"

type Achievement struct {
	ID             int64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code           *string       `gorm:"type:varchar(50);uniqueIndex" json:"code"` // 로직 매핑용 코드 (ex. FIRST_LOOT)
	Name           *string       `gorm:"type:varchar(100)" json:"name"`
	Description    *string       `gorm:"type:varchar(255)" json:"description"`
	ConditionJSON  types.JSONMap `gorm:"column:conditionJson;type:json" json:"conditionJson"` // 달성 조건 메타데이터
	BadgeImageName *string       `gorm:"column:badgeImageName;type:varchar(255)" json:"badgeImageName"`
	IsHidden       bool          `gorm:"column:isHidden;type:boolean;default:false" json:"isHidden"`

	// Associations
	UserAchievements []UserAchievement `gorm:"foreignKey:AchievementID" json:"userAchievements,omitempty"`
}

func (Achievement) TableName() string {
	return "achievements"
}
