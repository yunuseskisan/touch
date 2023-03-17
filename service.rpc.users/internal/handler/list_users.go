package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.users/internal/db"
	usersproto "github.com/yunuseskisan/touch/service.rpc.users/proto"
)

func ValidateListUsersRequest(req *usersproto.ListUsersRequest) error {
	if req.GetLimit() > 100 {
		return errors.New("limit must be no more than 100")
	}
	return nil
}

func (h Handler) ListUsers(ctx context.Context, req *usersproto.ListUsersRequest) (*usersproto.ListUsersResponse, error) {
	if err := ValidateListUsersRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &usersproto.ListUsersResponse{
		Users: []*usersproto.User{},
	}
	if err := func() error {
		users, err := db.SelectUsers(ctx, h.conn.Conn(), &db.SelectUsersQuery{
			Limit:  &req.Limit,
			Offset: &req.Offset,
		})
		if err != nil {
			return err
		}

		for _, user := range users {
			resp.Users = append(resp.Users, user.ToProto())
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
