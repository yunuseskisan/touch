package db

import (
	"context"

	"github.com/go-pg/pg/v10/orm"

	"github.com/yunuseskisan/touch/service.rpc.users/internal/models"
)

type SelectUsersQuery struct {
	Limit  *uint64
	Offset *uint64

	IDs []string
}

func (s SelectUsersQuery) Apply(q *orm.Query) *orm.Query {
	if s.Limit != nil {
		q.Limit(int(*s.Limit))
	}
	if s.Offset != nil {
		q.Offset(int(*s.Offset))
	}

	if len(s.IDs) > 0 {
		q.WhereIn(`"user"."id" IN (?)`, s.IDs)
	}
	return q
}

func SelectUsers(ctx context.Context, txn orm.DB, query *SelectUsersQuery) ([]*models.User, error) {
	var users []*models.User
	if err := query.
		Apply(txn.Model(&users)).
		Context(ctx).
		Select(); err != nil {
		return nil, err
	}
	return users, nil
}
