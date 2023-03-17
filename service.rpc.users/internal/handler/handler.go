package handler

import (
	"github.com/go-pg/pg/v10"

	usersproto "github.com/yunuseskisan/touch/service.rpc.users/proto"
)

type Handler struct {
	conn *pg.DB

	usersproto.UnimplementedUsersServiceServer
}

func New(conn *pg.DB) *Handler {
	return &Handler{
		conn: conn,
	}
}
