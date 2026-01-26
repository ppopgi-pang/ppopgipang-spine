package entities

import "time"

type UserSearchHistory struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     *uint     `gorm:"column:userId;index:idx_search_user_time" json:"userId"`
	Keyword    *string   `gorm:"type:varchar(100)" json:"keyword"`
	SearchedAt time.Time `gorm:"column:searchedAt;type:datetime(6);autoCreateTime" json:"searchedAt"`

	// Associations
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (UserSearchHistory) TableName() string {
	return "user_search_history"
}
