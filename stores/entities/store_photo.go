package entities

type StorePhoto struct {
	ID        int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID   *int64  `gorm:"column:storeId;index:idx_store_photo_type" json:"storeId"`
	Type      string  `gorm:"type:enum('cover','sign','inside','roadview')" json:"type"`
	ImageName *string `gorm:"column:imageName;type:varchar(255)" json:"imageName"`

	// Associations
	Store *Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
}

func (StorePhoto) TableName() string {
	return "store_photos"
}
