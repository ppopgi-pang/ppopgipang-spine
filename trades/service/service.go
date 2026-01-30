package service

import "gorm.io/gorm"

type TradeService struct {
	db *gorm.DB
}

func NewTradeService(db *gorm.DB) *TradeService {
	return &TradeService{db: db}
}
