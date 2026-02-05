package entities

type StoreOpeningHour struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID   *uint  `gorm:"column:storeId" json:"storeId"`
	DayOfWeek *int8  `gorm:"column:dayOfWeek;type:tinyint" json:"dayOfWeek"` // 0=Sun, 1=Mon, ... 6=Sat
	OpenTime  string `gorm:"column:openTime;type:time" json:"openTime"`
	CloseTime string `gorm:"column:closeTime;type:time" json:"closeTime"`
	IsClosed  int8   `gorm:"column:isClosed;type:tinyint;default:0" json:"isClosed"`

	// Associations
	Store *Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
}

func (StoreOpeningHour) TableName() string {
	return "store_opening_hours"
}
