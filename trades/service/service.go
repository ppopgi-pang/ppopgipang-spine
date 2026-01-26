package service

import "github.com/ppopgi-pang/ppopgipang-spine/trades/repository"

type TradeService struct {
	repo *repository.TradeRepository
}

func NewTradeService(repo *repository.TradeRepository) *TradeService {
	return &TradeService{repo: repo}
}
