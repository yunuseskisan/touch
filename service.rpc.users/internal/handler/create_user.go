package handler

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yunuseskisan/touch/service.rpc.users/internal/db"
	"github.com/yunuseskisan/touch/service.rpc.users/internal/models"
	usersproto "github.com/yunuseskisan/touch/service.rpc.users/proto"
)

func ValidateCreateUserRequest(req *usersproto.CreateUserRequest) error {
	if len(req.GetRequestId()) == 0 {
		return errors.New("request_id must be provided")
	}
	if !govalidator.IsUUID(req.GetRequestId()) {
		return errors.New("request_id must be a uuid")
	}
	if req.GetUser().GetFirstName() == "" {
		return errors.New("user.first_name must be provided")
	}
	if req.GetUser().GetLastName() == "" {
		return errors.New("user.last_name must be provided")
	}
	return nil
}

func (h Handler) CreateUser(ctx context.Context, req *usersproto.CreateUserRequest) (*usersproto.User, error) {
	if err := ValidateCreateUserRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := models.NewUser(req.GetUser())

	if err := h.conn.RunInTransaction(ctx, func(tx *pg.Tx) error {
		if err := db.InsertUser(ctx, tx, user); err != nil {
			return err
		}
		return nil
	}); err != nil {
		switch {
		case errors.Is(err, models.UserAlreadyExistsError):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Ops something went wrong!")
		}
	}

	return user.ToProto(), nil
}
