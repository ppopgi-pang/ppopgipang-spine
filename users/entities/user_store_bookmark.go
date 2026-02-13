package entities

import "time"

type UserStoreBookmark struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"column:userId;not null;" json:"userId"`
	StoreID   int64     `gorm:"column:storeId;not null;index:idx_user_store_bookmarks_store;uniqueIndex:uq_user_store_bookmarks_user_store" json:"storeId"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (UserStoreBookmark) TableName() string {
	return "user_store_bookmarks"
}
