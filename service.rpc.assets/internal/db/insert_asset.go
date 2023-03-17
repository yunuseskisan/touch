package db

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/models"
)

func convertAssetSQLError(err error) error {
	pqErr, ok := err.(pg.Error)
	if !ok {
		return err
	}
	switch pqErr.Field('n') {
	case "assets_pkey", "assets_user_id_type_key":
		return models.AssetAlreadyExistsError
	case "assets_user_id_fkey":
		return models.UserNotFoundError
	}
	return err
}

func InsertAsset(ctx context.Context, txn orm.DB, asset *models.Asset) error {
	if _, err := txn.
		Model(asset).
		Context(ctx).
		Insert(); err != nil {
		return convertAssetSQLError(err)
	}
	return nil
}
