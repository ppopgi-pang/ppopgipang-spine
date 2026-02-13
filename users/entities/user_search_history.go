package entities

import "time"

type UserSearchHistory struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     *int64    `gorm:"column:userId;index:idx_search_user_time" json:"userId"`
	Keyword    *string   `gorm:"type:varchar(100)" json:"keyword"`
	SearchedAt time.Time `gorm:"column:searchedAt;type:datetime(6);autoCreateTime" json:"searchedAt"`

	// Associations
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (UserSearchHistory) TableName() string {
	return "user_search_history"
}
