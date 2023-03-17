package test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
	usersproto "github.com/yunuseskisan/touch/service.rpc.users/proto"
)

func TestTouch(t *testing.T) {
	ctx := context.Background()

	usersConn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		require.NoError(t, err)
	}
	usersClient := usersproto.NewUsersServiceClient(usersConn)

	assetsConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		require.NoError(t, err)
	}
	assetsClient := assetsproto.NewAssetsServiceClient(assetsConn)

	ordersConn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		require.NoError(t, err)
	}
	ordersClient := ordersproto.NewOrdersServiceClient(ordersConn)

	var (
		userID  string
		assetID string
	)

	t.Run("should successfully creates a user", func(t *testing.T) {
		user, err := usersClient.CreateUser(ctx, &usersproto.CreateUserRequest{
			RequestId: uuid.NewV4().String(),
			User: &usersproto.User{
				FirstName: "Yunus",
				LastName:  "Eskisan",
			},
		})

		require.NoError(t, err)
		assert.NotEmpty(t, user.GetId())

		userID = user.GetId()
	})

	t.Run("should successfully creates an asset for user", func(t *testing.T) {
		asset, err := assetsClient.CreateAsset(ctx, &assetsproto.CreateAssetRequest{
			RequestId: uuid.NewV4().String(),
			Asset: &assetsproto.Asset{
				UserId: userID,
				Type:   assetsproto.AssetType_ASSET_TYPE_GB00B3X7QG63,
			},
		})

		require.NoError(t, err)
		assert.NotEmpty(t, asset.GetId())

		assetID = asset.GetId()
	})

	t.Run("should successfully buy units for asset", func(t *testing.T) {
		asset, err := assetsClient.BuyAsset(ctx, &assetsproto.BuyAssetRequest{
			RequestId: uuid.NewV4().String(),
			AssetId:   assetID,
			Units:     10,
		})

		require.NoError(t, err)
		assert.Equal(t, int32(10), asset.GetUnits())
	})

	t.Run("should have successfully created buy order", func(t *testing.T) {
		orders, err := ordersClient.ListOrders(ctx, &ordersproto.ListOrdersRequest{
			Limit: 10,
		})

		require.NoError(t, err)
		assert.Len(t, orders.GetOrders(), 1)
		assert.Equal(t, int32(10), orders.GetOrders()[0].GetAmount())
		assert.Equal(t, ordersproto.Currency_CURRENCY_UNITS, orders.GetOrders()[0].GetCurrency())
		assert.Equal(t, ordersproto.Instruction_INSTRUCTION_BUY, orders.GetOrders()[0].GetInstruction())
		assert.Equal(t, assetsproto.AssetType_ASSET_TYPE_GB00B3X7QG63, orders.GetOrders()[0].GetAssetType())
	})

	t.Run("should successfully buy further units for asset", func(t *testing.T) {
		asset, err := assetsClient.BuyAsset(ctx, &assetsproto.BuyAssetRequest{
			RequestId: uuid.NewV4().String(),
			AssetId:   assetID,
			Units:     5,
		})

		require.NoError(t, err)
		assert.Equal(t, int32(15), asset.GetUnits())
	})

	t.Run("should have successfully created buy order", func(t *testing.T) {
		orders, err := ordersClient.ListOrders(ctx, &ordersproto.ListOrdersRequest{
			Limit: 10,
		})

		require.NoError(t, err)
		assert.Len(t, orders.GetOrders(), 2)
		assert.Equal(t, int32(5), orders.GetOrders()[1].GetAmount())
		assert.Equal(t, ordersproto.Currency_CURRENCY_UNITS, orders.GetOrders()[1].GetCurrency())
		assert.Equal(t, ordersproto.Instruction_INSTRUCTION_BUY, orders.GetOrders()[1].GetInstruction())
		assert.Equal(t, assetsproto.AssetType_ASSET_TYPE_GB00B3X7QG63, orders.GetOrders()[1].GetAssetType())
	})

	t.Run("should successfully sell units for asset", func(t *testing.T) {
		asset, err := assetsClient.SellAsset(ctx, &assetsproto.SellAssetRequest{
			RequestId: uuid.NewV4().String(),
			AssetId:   assetID,
			Units:     15,
		})

		require.NoError(t, err)
		assert.Equal(t, int32(0), asset.GetUnits())
	})

	t.Run("should have successfully created sell order", func(t *testing.T) {
		orders, err := ordersClient.ListOrders(ctx, &ordersproto.ListOrdersRequest{
			Limit: 10,
		})

		require.NoError(t, err)
		assert.Len(t, orders.GetOrders(), 3)
		assert.Equal(t, int32(15), orders.GetOrders()[2].GetAmount())
		assert.Equal(t, ordersproto.Currency_CURRENCY_UNITS, orders.GetOrders()[2].GetCurrency())
		assert.Equal(t, ordersproto.Instruction_INSTRUCTION_SELL, orders.GetOrders()[2].GetInstruction())
		assert.Equal(t, assetsproto.AssetType_ASSET_TYPE_GB00B3X7QG63, orders.GetOrders()[2].GetAssetType())
	})
}
