package db

import (
	"context"

	"github.com/go-pg/pg/v10/orm"

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/models"
)

type SelectAssetsQuery struct {
	Limit  *uint64
	Offset *uint64

	IDs     []string
	UserIDs []string
}

func (s SelectAssetsQuery) Apply(q *orm.Query) *orm.Query {
	if s.Limit != nil {
		q.Limit(int(*s.Limit))
	}
	if s.Offset != nil {
		q.Offset(int(*s.Offset))
	}

	if len(s.IDs) > 0 {
		q.WhereIn(`"asset"."id" IN (?)`, s.IDs)
	}
	if len(s.UserIDs) > 0 {
		q.WhereIn(`"asset"."user_id" IN (?)`, s.UserIDs)
	}
	return q
}

func SelectAssets(ctx context.Context, txn orm.DB, query *SelectAssetsQuery) ([]*models.Asset, error) {
	var assets []*models.Asset
	if err := query.
		Apply(txn.Model(&assets)).
		Context(ctx).
		Select(); err != nil {
		return nil, err
	}
	return assets, nil
}
