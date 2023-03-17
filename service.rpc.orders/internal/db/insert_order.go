package db

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/yunuseskisan/touch/service.rpc.orders/internal/models"
)

func convertOrderSQLError(err error) error {
	pqErr, ok := err.(pg.Error)
	if !ok {
		return err
	}
	switch pqErr.Field('n') {
	case "orders_pkey":
		return models.OrderAlreadyExistsError
	}
	return err
}

func InsertOrder(ctx context.Context, txn orm.DB, order *models.Order) error {
	if _, err := txn.
		Model(order).
		Context(ctx).
		Insert(); err != nil {
		return convertOrderSQLError(err)
	}
	return nil
}
