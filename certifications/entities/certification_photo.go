package entities

type CertificationPhoto struct {
	ID              int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	CertificationID *int64  `gorm:"column:certificationId" json:"certificationId"`
	ImageName       *string `gorm:"column:imageName;type:varchar(255)" json:"imageName"`
	SortOrder       int     `gorm:"column:sortOrder;default:0" json:"sortOrder"`

	// Associations
	Certification *Certification `gorm:"foreignKey:CertificationID;constraint:OnDelete:CASCADE" json:"certification,omitempty"`
}

func (CertificationPhoto) TableName() string {
	return "certification_photos"
}
