package migrations

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

var Collection = migrations.NewCollection().SetTableName("orders_migrations")

func init() {
	err := Collection.DiscoverSQLMigrations("service.rpc.orders/internal/db/migrations")
	if err != nil {
		log.Fatal(err)
	}
}
