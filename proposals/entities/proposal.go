package entities

import (
	"time"

	stores "github.com/ppopgi-pang/ppopgipang-spine/stores/entities"
	users "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
)

type Proposal struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      *string   `gorm:"type:varchar(100)" json:"name"`
	Address   *string   `gorm:"type:varchar(255)" json:"address"`
	Latitude  *float64  `gorm:"type:decimal(10,6)" json:"latitude"`
	Longitude *float64  `gorm:"type:decimal(10,6)" json:"longitude"`
	Status    string    `gorm:"type:enum('pending','approved','rejected');default:pending" json:"status"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UserID    *int64    `gorm:"column:userId" json:"userId"`
	StoreID   *int64    `gorm:"column:storeId" json:"storeId"` // 승인 시 연결된 store ID

	// Certification
	User  *users.User   `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" json:"user,omitempty"`
	Store *stores.Store `gorm:"foreignKey:StoreID;constraint:OnDelete:SET NULL" json:"store,omitempty"`
}

func (Proposal) TableName() string {
	return "proposals"
}
