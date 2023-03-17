package db

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/yunuseskisan/touch/service.rpc.users/internal/models"
)

func convertUserSQLError(err error) error {
	pqErr, ok := err.(pg.Error)
	if !ok {
		return err
	}
	switch pqErr.Field('n') {
	case "users_pkey":
		return models.UserAlreadyExistsError
	}
	return err
}

func InsertUser(ctx context.Context, txn orm.DB, user *models.User) error {
	if _, err := txn.
		Model(user).
		Context(ctx).
		Insert(); err != nil {
		return convertUserSQLError(err)
	}
	return nil
}
