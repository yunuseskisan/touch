package handler

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
	"github.com/yunuseskisan/touch/service.rpc.orders/internal/db"
	"github.com/yunuseskisan/touch/service.rpc.orders/internal/models"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

func ValidateCreateOrderRequest(req *ordersproto.CreateOrderRequest) error {
	if len(req.GetRequestId()) == 0 {
		return errors.New("request_id must be provided")
	}
	if !govalidator.IsUUID(req.GetRequestId()) {
		return errors.New("request_id must be a uuid")
	}
	if req.GetOrder().GetAssetType() == assetsproto.AssetType_ASSET_TYPE_UNKNOWN {
		return errors.New("order.asset_type must be provided")
	}
	if req.GetOrder().GetInstruction() == ordersproto.Instruction_INSTRUCTION_UNKNOWN {
		return errors.New("order.instruction must be provided")
	}
	if req.GetOrder().GetAmount() == 0 {
		return errors.New("order.amount must be provided")
	}
	if req.GetOrder().GetCurrency() == ordersproto.Currency_CURRENCY_UNKNOWN {
		return errors.New("order.currency must be provided")
	}
	return nil
}

func (h Handler) CreateOrder(ctx context.Context, req *ordersproto.CreateOrderRequest) (*ordersproto.Order, error) {
	if err := ValidateCreateOrderRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	order := models.NewOrder(req.GetOrder())

	if err := h.conn.RunInTransaction(ctx, func(tx *pg.Tx) error {
		if err := db.InsertOrder(ctx, tx, order); err != nil {
			return err
		}
		return nil
	}); err != nil {
		switch {
		case errors.Is(err, models.OrderAlreadyExistsError):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}

	return order.ToProto(), nil
}
