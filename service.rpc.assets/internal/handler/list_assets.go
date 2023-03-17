package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/db"
	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
)

func ValidateListAssetsRequest(req *assetsproto.ListAssetsRequest) error {
	if req.GetLimit() > 100 {
		return errors.New("limit must be no more than 100")
	}
	for i, id := range req.GetUserIds() {
		if !govalidator.IsUUID(id) {
			return fmt.Errorf("user[%d] must be a uuid", i)
		}
	}
	return nil
}

func (h Handler) ListAssets(ctx context.Context, req *assetsproto.ListAssetsRequest) (*assetsproto.ListAssetsResponse, error) {
	if err := ValidateListAssetsRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &assetsproto.ListAssetsResponse{
		Assets: []*assetsproto.Asset{},
	}
	if err := func() error {
		assets, err := db.SelectAssets(ctx, h.conn.Conn(), &db.SelectAssetsQuery{
			Limit:   &req.Limit,
			Offset:  &req.Offset,
			UserIDs: req.GetUserIds(),
		})
		if err != nil {
			return err
		}

		for _, asset := range assets {
			resp.Assets = append(resp.Assets, asset.ToProto())
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
