package entities

type CertificationTag struct {
	CertificationID int64 `gorm:"column:certificationId;primaryKey" json:"certificationId"`
	TagID           int64 `gorm:"column:tagId;primaryKey" json:"tagId"`
}

func (CertificationTag) TableName() string {
	return "certification_tags"
}
