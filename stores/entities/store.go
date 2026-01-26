package entities

import "time"

type Store struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Address       *string   `gorm:"type:varchar(255)" json:"address"`
	Region1       *string   `gorm:"type:varchar(50)" json:"region1"` // 시/도 (예: 서울특별시)
	Region2       *string   `gorm:"type:varchar(50)" json:"region2"` // 구/군 (예: 마포구)
	Latitude      float64   `gorm:"type:decimal(10,6);not null" json:"latitude"`
	Longitude     float64   `gorm:"type:decimal(10,6);not null" json:"longitude"`
	Phone         *string   `gorm:"type:varchar(20)" json:"phone"`
	AverageRating float32   `gorm:"column:averageRating;type:float;default:0" json:"averageRating"`
	TypeID        *uint     `gorm:"column:typeId" json:"typeId"`
	CreatedAt     time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`

	// Associations
	Type           *StoreType                     `gorm:"foreignKey:TypeID;constraint:OnDelete:SET NULL" json:"type,omitempty"`
	Analytics      *StoreAnalytics                `gorm:"foreignKey:StoreID" json:"analytics,omitempty"`
	Photos         []StorePhoto                   `gorm:"foreignKey:StoreID" json:"photos,omitempty"`
	OpeningHours   []StoreOpeningHour             `gorm:"foreignKey:StoreID" json:"openingHours,omitempty"`
	Facilities     *StoreFacility                 `gorm:"foreignKey:StoreID" json:"facilities,omitempty"`
}

func (Store) TableName() string {
	return "stores"
}
