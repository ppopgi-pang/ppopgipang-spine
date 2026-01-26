package entities

type CertificationTag struct {
	CertificationID uint `gorm:"column:certificationId;primaryKey" json:"certificationId"`
	TagID           uint `gorm:"column:tagId;primaryKey" json:"tagId"`
}

func (CertificationTag) TableName() string {
	return "certification_tags"
}
