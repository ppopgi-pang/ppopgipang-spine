package entities

import (
	"time"

	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type LootLike struct {
	UserID          int64     `gorm:"column:userId;primaryKey" json:"userId"`
	CertificationID int64     `gorm:"column:certificationId;primaryKey" json:"certificationId"`
	CreatedAt       time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	User          *users.User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Certification *Certification `gorm:"foreignKey:CertificationID;constraint:OnDelete:CASCADE" json:"certification,omitempty"`
}

func (LootLike) TableName() string {
	return "loot_likes"
}
