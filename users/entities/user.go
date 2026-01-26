package entities

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Email         string    `gorm:"type:varchar(50);not null;uniqueIndex:uq_users_email" json:"email"`
	KakaoID       *string   `gorm:"column:kakaoId;type:varchar(255);uniqueIndex:uq_users_kakao" json:"kakaoId"`
	Nickname      string    `gorm:"type:varchar(30);not null" json:"nickname"`
	ProfileImage  *string   `gorm:"column:profileImage;type:varchar(255)" json:"profileImage"`
	IsAdmin       int8      `gorm:"column:isAdmin;type:tinyint;not null;default:0" json:"isAdmin"`
	RefreshToken  *string   `gorm:"column:refreshToken;type:varchar(255)" json:"refreshToken"`
	AdminPassword *string   `gorm:"column:adminPassword;type:varchar(255)" json:"adminPassword"`
	MannerTemp    float64   `gorm:"column:mannerTemp;type:decimal(4,1);default:36.5" json:"mannerTemp"` // 매너 온도
	CreatedAt     time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`

	// Associations
	UserProgress     *UserProgress                   `gorm:"foreignKey:UserID" json:"userProgress,omitempty"`
	UserLoots        []UserLoot                      `gorm:"foreignKey:UserID" json:"userLoots,omitempty"`
}

func (User) TableName() string {
	return "users"
}
