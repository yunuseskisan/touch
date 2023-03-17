package migrations

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

var Collection = migrations.NewCollection().SetTableName("users_migrations")

func init() {
	err := Collection.DiscoverSQLMigrations("service.rpc.users/internal/db/migrations")
	if err != nil {
		log.Fatal(err)
	}
}
