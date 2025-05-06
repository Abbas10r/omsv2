package main

import (
	"github.com/abbas10r/common"
	"github.com/abbas10r/common/api"
	_ "github.com/abbas10r/common/api"
	"net/http"
)

type handler struct {
	client api.OrderServiceClient
}

func NewHandler(client api.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*api.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(r.Context(), &api.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
