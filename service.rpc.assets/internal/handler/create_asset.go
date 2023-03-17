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
)

func ValidateCreateAssetRequest(req *assetsproto.CreateAssetRequest) error {
	if len(req.GetRequestId()) == 0 {
		return errors.New("request_id must be provided")
	}
	if !govalidator.IsUUID(req.GetRequestId()) {
		return errors.New("request_id must be a uuid")
	}
	if req.GetAsset().GetType() == assetsproto.AssetType_ASSET_TYPE_UNKNOWN {
		return errors.New("asset.type must be provided")
	}
	return nil
}

func (h Handler) CreateAsset(ctx context.Context, req *assetsproto.CreateAssetRequest) (*assetsproto.Asset, error) {
	if err := ValidateCreateAssetRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	asset := models.NewAsset(req.GetAsset())

	if err := h.conn.RunInTransaction(ctx, func(tx *pg.Tx) error {
		if err := db.InsertAsset(ctx, tx, asset); err != nil {
			return err
		}
		return nil
	}); err != nil {
		switch {
		case errors.Is(err, models.AssetAlreadyExistsError):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case errors.Is(err, models.UserNotFoundError):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}

	return asset.ToProto(), nil
}
