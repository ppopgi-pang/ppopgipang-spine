package routes

import (
	"github.com/NARUBROWN/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/trades/controller"
)

func RegisterTradeRoutes(app spine.App) {
	app.Route("POST", "/v1/trades", (*controller.TradeController).CreateTrade)
}
