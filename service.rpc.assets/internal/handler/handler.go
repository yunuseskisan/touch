package handler

import (
	"github.com/go-pg/pg/v10"

	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

type Handler struct {
	conn         *pg.DB
	ordersClient ordersproto.OrdersServiceClient

	assetsproto.UnimplementedAssetsServiceServer
}

func New(conn *pg.DB, ordersClient ordersproto.OrdersServiceClient) *Handler {
	return &Handler{
		conn:         conn,
		ordersClient: ordersClient,
	}
}
