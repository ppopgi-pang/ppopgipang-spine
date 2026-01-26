package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/service"
)

type StoreController struct {
	service *service.StoreService
}

func NewStoreController(storesService *service.StoreService) *StoreController {
	return &StoreController{service: storesService}
}

func (s *StoreController) FindNearestStore(ctx context.Context, query query.Values) {}
