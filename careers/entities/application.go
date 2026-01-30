package entities

import "time"

type Application struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	JobPostingID int64     `gorm:"column:jobPosting_id;not null;index:idx_applications_job_posting" json:"jobPostingId"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Email        string    `gorm:"type:varchar(200);not null" json:"email"`
	Phone        *string   `gorm:"type:varchar(20)" json:"phone"`
	ResumeName   *string   `gorm:"column:resumeName;type:varchar(500)" json:"resumeName"`
	Memo         *string   `gorm:"type:text" json:"memo"`
	Status       string    `gorm:"type:varchar(20);default:new;index:idx_applications_status" json:"status"`
	CreatedAt    time.Time `gorm:"column:createdAt;type:datetime(6);autoCreateTime" json:"createdAt"`

	// Associations
	JobPosting *JobPosting `gorm:"foreignKey:JobPostingID;constraint:OnDelete:CASCADE" json:"jobPosting,omitempty"`
}

func (Application) TableName() string {
	return "applications"
}
