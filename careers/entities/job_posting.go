package entities

import "time"

type JobPosting struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title        string    `gorm:"type:varchar(200);not null" json:"title"`
	Description  *string   `gorm:"type:text" json:"description"`
	Department   *string   `gorm:"type:varchar(100)" json:"department"`
	PositionType *string   `gorm:"column:positionType;type:varchar(50)" json:"positionType"`
	Location     *string   `gorm:"type:varchar(200)" json:"location"`
	IsActive     bool      `gorm:"column:isActive;type:boolean;default:true" json:"isActive"`
	CreatedAt    time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`

	// Associations
	Applications []Application `gorm:"foreignKey:JobPostingID" json:"applications,omitempty"`
}

func (JobPosting) TableName() string {
	return "job_postings"
}
