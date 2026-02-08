package entities

import (
	"time"

	stores "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type Certification struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int64     `gorm:"column:userId;not null;index:idx_cert_user_time" json:"userId"`
	StoreID    int64     `gorm:"column:storeId;not null;index:idx_cert_store_time" json:"storeId"`
	Type       string    `gorm:"type:enum('loot','checkin');not null" json:"type"`              // loot:득템, checkin:빈손방문
	OccurredAt time.Time `gorm:"column:occurredAt;type:datetime(6);not null" json:"occurredAt"` // 실제 방문 시간
	Latitude   *float64  `gorm:"type:decimal(10,6)" json:"latitude"`                            // 인증 당시 GPS
	Longitude  *float64  `gorm:"type:decimal(10,6)" json:"longitude"`
	Exp        int       `gorm:"not null" json:"exp"`                            // 획득 경험치
	Comment    *string   `gorm:"type:varchar(200)" json:"comment"`               // 한줄평 (득템용)
	Rating     *string   `gorm:"type:enum('good','normal','bad')" json:"rating"` // 상태 평가 (체크인용)
	CreatedAt  time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	User    *users.User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Store   *stores.Store         `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
	Photos  []CertificationPhoto  `gorm:"foreignKey:CertificationID" json:"photos,omitempty"`
	Tags    []users.LootTag       `gorm:"many2many:certification_tags" json:"tags,omitempty"`
	Reasons []CheckinReasonPreset `gorm:"many2many:certification_reasons" json:"reasons,omitempty"`
	Likes   []LootLike            `gorm:"foreignKey:CertificationID" json:"likes,omitempty"`
}

func (Certification) TableName() string {
	return "certifications"
}
