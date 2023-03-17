package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.orders/internal/db"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

func ValidateListOrdersRequest(req *ordersproto.ListOrdersRequest) error {
	if req.GetLimit() > 100 {
		return errors.New("limit must be no more than 100")
	}
	return nil
}

func (h Handler) ListOrders(ctx context.Context, req *ordersproto.ListOrdersRequest) (*ordersproto.ListOrdersResponse, error) {
	if err := ValidateListOrdersRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &ordersproto.ListOrdersResponse{
		Orders: []*ordersproto.Order{},
	}
	if err := func() error {
		orders, err := db.SelectOrders(ctx, h.conn.Conn(), &db.SelectOrdersQuery{
			Limit:  &req.Limit,
			Offset: &req.Offset,
		})
		if err != nil {
			return err
		}

		for _, order := range orders {
			resp.Orders = append(resp.Orders, order.ToProto())
		}
		return nil
	}(); err != nil {
		switch {
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}
	return resp, nil
}
