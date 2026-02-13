package entities

type CertificationReason struct {
	CertificationID int64 `gorm:"column:certificationId;primaryKey" json:"certificationId"`
	ReasonID        int64 `gorm:"column:reasonId;primaryKey" json:"reasonId"`
}

func (CertificationReason) TableName() string {
	return "certification_reasons"
}
