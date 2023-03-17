package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

var OrderAlreadyExistsError = errors.New("order already exists")
var OrderNotFoundError = errors.New("order not found")

type Order struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time

	AssetType   assetsproto.AssetType
	Instruction ordersproto.Instruction
	Amount      int32
	Currency    ordersproto.Currency
}

func NewOrder(pb *ordersproto.Order) *Order {
	return &Order{
		Id:          uuid.NewV4().String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		AssetType:   pb.GetAssetType(),
		Instruction: pb.Instruction,
		Amount:      pb.Amount,
		Currency:    pb.Currency,
	}
}

func (m Order) ToProto() *ordersproto.Order {
	return &ordersproto.Order{
		Id:          m.Id,
		CreatedAt:   timestamppb.New(m.CreatedAt),
		UpdatedAt:   timestamppb.New(m.UpdatedAt),
		AssetType:   m.AssetType,
		Instruction: m.Instruction,
		Amount:      m.Amount,
		Currency:    m.Currency,
	}
}
