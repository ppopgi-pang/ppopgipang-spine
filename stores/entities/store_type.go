package entities

type StoreType struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"type:varchar(50);not null" json:"name"` // 예: 인형뽑기, 가챠샵, 오락실
	Description *string `gorm:"type:varchar(255)" json:"description"`

	// Associations
	Stores []Store `gorm:"foreignKey:TypeID" json:"stores,omitempty"`
}

func (StoreType) TableName() string {
	return "store_type"
}
