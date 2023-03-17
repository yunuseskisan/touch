package migrations

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

var Collection = migrations.NewCollection().SetTableName("assets_migrations")

func init() {
	err := Collection.DiscoverSQLMigrations("service.rpc.assets/internal/db/migrations")
	if err != nil {
		log.Fatal(err)
	}
}
