package main

import (
	"context"
	"github.com/abbas10r/common"
	"github.com/abbas10r/common/api"
	"log"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, p *api.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}
	mergedItems := mergeItemsQuantities(p.Items)
	log.Print(mergedItems)

	//validate with stock service later

	return nil
}

func mergeItemsQuantities(items []*api.ItemsWithQuantity) []*api.ItemsWithQuantity {
	merged := make([]*api.ItemsWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
