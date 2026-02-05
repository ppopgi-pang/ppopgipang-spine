package entities

import (
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/commons/types"
)

type StoreAnalytics struct {
	StoreID         int64         `gorm:"column:storeId;primaryKey" json:"storeId"`
	CongestionScore int           `gorm:"column:congestionScore;default:0" json:"congestionScore"`      // 혼잡도 점수 (0~100)
	SuccessProb     int           `gorm:"column:successProb;default:50" json:"successProb"`             // AI 예측 득템 확률 (0~100)
	RecentLootCount int           `gorm:"column:recentLootCount;default:0" json:"recentLootCount"`      // 최근 1시간 내 득템 인증 수
	HotTimeJSON     types.JSONMap `gorm:"column:hotTimeJson;type:json" json:"hotTimeJson"`              // 시간대별 성공 확률 그래프 데이터
	LastAnalyzedAt  *time.Time    `gorm:"column:lastAnalyzedAt;type:datetime(6)" json:"lastAnalyzedAt"` // 마지막 분석 시각

	// Associations
	Store *Store `gorm:"foreignKey:StoreID;constraint:OnDelete:CASCADE" json:"store,omitempty"`
}

func (StoreAnalytics) TableName() string {
	return "store_analytics"
}
