package handler

import (
	"github.com/go-pg/pg/v10"

	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

type Handler struct {
	conn *pg.DB

	ordersproto.UnimplementedOrdersServiceServer
}

func New(conn *pg.DB) *Handler {
	return &Handler{
		conn: conn,
	}
}
