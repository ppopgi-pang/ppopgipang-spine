package controller

import (
	"context"

	"github.com/ppopgi-pang/ppopgipang-spine/trades/service"
)

type TradeController struct {
	service *service.TradeService
}

func NewTradeController(tradesService *service.TradeService) *TradeController {
	return &TradeController{service: tradesService}
}

func (t *TradeController) CreateTrade(ctx context.Context) {}
