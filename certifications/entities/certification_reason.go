package entities

type CertificationReason struct {
	CertificationID uint `gorm:"column:certificationId;primaryKey" json:"certificationId"`
	ReasonID        uint `gorm:"column:reasonId;primaryKey" json:"reasonId"`
}

func (CertificationReason) TableName() string {
	return "certification_reasons"
}
