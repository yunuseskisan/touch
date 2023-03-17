package handler

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/db"
	"github.com/yunuseskisan/touch/service.rpc.assets/internal/models"
	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

func ValidateBuyAssetRequest(req *assetsproto.BuyAssetRequest) error {
	if len(req.GetRequestId()) == 0 {
		return errors.New("request_id must be provided")
	}
	if !govalidator.IsUUID(req.GetRequestId()) {
		return errors.New("request_id must be a uuid")
	}
	if len(req.GetAssetId()) == 0 {
		return errors.New("asset_id must be provided")
	}
	if req.GetUnits() == 0 {
		return errors.New("uints must be provided")
	}
	return nil
}

func (h Handler) BuyAsset(ctx context.Context, req *assetsproto.BuyAssetRequest) (*assetsproto.Asset, error) {
	if err := ValidateBuyAssetRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &assetsproto.Asset{}
	if err := h.conn.RunInTransaction(ctx, func(tx *pg.Tx) error {
		assets, err := db.SelectAssets(ctx, tx, &db.SelectAssetsQuery{
			IDs: []string{req.GetAssetId()},
		})

		if err != nil {
			return err
		}

		if len(assets) == 0 {
			return models.AssetNotFoundError
		}

		asset := assets[0]
		asset.Units += req.Units

		if err := db.UpdateAsset(ctx, tx, asset, "units"); err != nil {
			return err
		}

		// Perform this operation async to guarantee success.
		h.ordersClient.CreateOrder(ctx, &ordersproto.CreateOrderRequest{
			RequestId: req.GetRequestId(),
			Order: &ordersproto.Order{
				AssetType:   asset.Type,
				Instruction: ordersproto.Instruction_INSTRUCTION_BUY,
				Amount:      req.Units,
				Currency:    ordersproto.Currency_CURRENCY_UNITS,
			},
		})

		resp = assets[0].ToProto()
		return nil
	}); err != nil {
		switch {
		case errors.Is(err, models.AssetNotFoundError):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}
	return resp, nil
}
