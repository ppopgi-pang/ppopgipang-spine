package entities

import "github.com/ppopgi-pang/ppopgipang-spine/commons/types"

type StoreFacility struct {
	StoreID        int64                 `gorm:"column:storeId;primaryKey" json:"storeId"`
	MachineCount   *int                  `gorm:"column:machineCount" json:"machineCount"`
	PaymentMethods types.JSONStringArray `gorm:"column:paymentMethods;type:json" json:"paymentMethods"` // ["cash", "card", "qr"]
	Notes          *string               `gorm:"type:varchar(255)" json:"notes"`

	// Associations
	Store *Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
}

func (StoreFacility) TableName() string {
	return "store_facilities"
}
