package entities

import "time"

type UserLoot struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          *uint     `gorm:"column:userId" json:"userId"`
	CertificationID *uint     `gorm:"column:certificationId" json:"certificationId"` // 어떤 인증을 통해 획득했는지
	Title           *string   `gorm:"type:varchar(100)" json:"title"`                // 유저가 입력하거나 AI가 추천한 이름
	Category        *string   `gorm:"type:varchar(50)" json:"category"`              // 인형, 피규어, 키링 등
	EstimatedPrice  *int      `gorm:"column:estimatedPrice" json:"estimatedPrice"`   // 시장 추정가
	Rarity          string    `gorm:"type:enum('common','rare','legend');default:common" json:"rarity"`
	AIConfidence    float32   `gorm:"column:aiConfidence;type:float;default:0" json:"aiConfidence"` // AI Vision 인식 신뢰도 (0.0~1.0)
	Status          string    `gorm:"type:enum('kept','selling','sold','exchanged');default:kept" json:"status"`
	CreatedAt       time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (UserLoot) TableName() string {
	return "user_loots"
}
