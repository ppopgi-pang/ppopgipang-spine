package dto

import "time"

type CheckInResponse struct {
	ID                 int64     `json:"id" example:"1001"`
	OccurredAt         time.Time `json:"occurred_at" example:"2026-03-23T12:00:00Z"`
	CertificationCount int64     `json:"certification_count" example:"5"`
}
