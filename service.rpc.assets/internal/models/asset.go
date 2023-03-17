package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
)

var AssetAlreadyExistsError = errors.New("asset already exists")

var AssetNotFoundError = errors.New("asset not found")

var UserNotFoundError = errors.New("user not found")

var InsufficientBalanceError = errors.New("asset does not have sufficient balance to complete sale")

type Asset struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID string
	Type   assetsproto.AssetType
	Units  int32
}

func NewAsset(pb *assetsproto.Asset) *Asset {
	return &Asset{
		Id:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    pb.GetUserId(),
		Type:      pb.GetType(),
		Units:     0,
	}
}

func (m Asset) ToProto() *assetsproto.Asset {
	return &assetsproto.Asset{
		Id:        m.Id,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
		UserId:    m.UserID,
		Type:      m.Type,
		Units:     m.Units,
	}
}
