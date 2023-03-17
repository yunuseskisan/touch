package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/db"
	"github.com/yunuseskisan/touch/service.rpc.assets/internal/models"
	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
)

func ValidateGetAssetRequest(req *assetsproto.GetAssetRequest) error {
	if len(req.GetId()) == 0 {
		return errors.New("id must be provided")
	}
	return nil
}

func (h Handler) GetAsset(ctx context.Context, req *assetsproto.GetAssetRequest) (*assetsproto.Asset, error) {
	if err := ValidateGetAssetRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &assetsproto.Asset{}
	if err := func() error {
		assets, err := db.SelectAssets(ctx, h.conn.Conn(), &db.SelectAssetsQuery{
			IDs: []string{req.GetId()},
		})
		if err != nil {
			return err
		}

		if len(assets) == 0 {
			return models.AssetNotFoundError
		}

		resp = assets[0].ToProto()
		return nil
	}(); err != nil {
		switch {
		case errors.Is(err, models.AssetNotFoundError):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}
	return resp, nil
}
