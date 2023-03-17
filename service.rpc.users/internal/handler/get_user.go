package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.users/internal/db"
	"github.com/yunuseskisan/touch/service.rpc.users/internal/models"
	usersproto "github.com/yunuseskisan/touch/service.rpc.users/proto"
)

func ValidateGetUserRequest(req *usersproto.GetUserRequest) error {
	if len(req.GetId()) == 0 {
		return errors.New("id must be provided")
	}
	return nil
}

func (h Handler) GetUser(ctx context.Context, req *usersproto.GetUserRequest) (*usersproto.User, error) {
	if err := ValidateGetUserRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &usersproto.User{}
	if err := func() error {
		users, err := db.SelectUsers(ctx, h.conn.Conn(), &db.SelectUsersQuery{
			IDs: []string{req.GetId()},
		})
		if err != nil {
			return err
		}

		if len(users) == 0 {
			return models.UserNotFoundError
		}

		resp = users[0].ToProto()
		return nil
	}(); err != nil {
		switch {
		case errors.Is(err, models.UserNotFoundError):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}
	return resp, nil
}
