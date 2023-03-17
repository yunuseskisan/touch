package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	usersproto "github.com/yunuseskisan/touch/service.rpc.users/proto"
)

var UserAlreadyExistsError = errors.New("user already exists")
var UserNotFoundError = errors.New("user not found")

type User struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time

	FirstName string
	LastName  string
}

func NewUser(pb *usersproto.User) *User {
	return &User{
		Id:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: pb.FirstName,
		LastName:  pb.LastName,
	}
}

func (m User) ToProto() *usersproto.User {
	return &usersproto.User{
		Id:        m.Id,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}
