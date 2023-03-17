package db

import (
	"context"

	"github.com/go-pg/pg/v10/orm"

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/models"
)

func UpdateAsset(ctx context.Context, txn orm.DB, asset *models.Asset, fields ...string) error {
	if _, err := txn.Model(asset).
		Context(ctx).
		Column(fields...).
		WherePK().
		Update(); err != nil {
		return err
	}
	return nil
}
