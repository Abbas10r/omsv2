package main

import (
	"context"
	"github.com/abbas10r/common/api"
	"google.golang.org/grpc"
	"log"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
	service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{
		service: service,
	}
	api.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *api.CreateOrderRequest) (*api.Order, error) {
	log.Println("New order received! Order %v", p)
	o := &api.Order{
		ID: "42",
	}
	return o, nil
}
