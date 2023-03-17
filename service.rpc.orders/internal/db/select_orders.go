package db

import (
	"context"

	"github.com/go-pg/pg/v10/orm"

	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
	"github.com/yunuseskisan/touch/service.rpc.orders/internal/models"
)

type SelectOrdersQuery struct {
	Limit  *uint64
	Offset *uint64

	AssetTypes []assetsproto.AssetType
}

func (s SelectOrdersQuery) Apply(q *orm.Query) *orm.Query {
	if s.Limit != nil {
		q.Limit(int(*s.Limit))
	}
	if s.Offset != nil {
		q.Offset(int(*s.Offset))
	}

	if len(s.AssetTypes) > 0 {
		q.WhereIn(`"order"."asset_type" IN (?)`, s.AssetTypes)
	}
	return q
}

func SelectOrders(ctx context.Context, txn orm.DB, query *SelectOrdersQuery) ([]*models.Order, error) {
	var orders []*models.Order
	if err := query.
		Apply(txn.Model(&orders)).
		Context(ctx).
		Select(); err != nil {
		return nil, err
	}
	return orders, nil
}
