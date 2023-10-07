package db

import (
	"context"
	"github.com/dipay/db/migrations"
	"github.com/dipay/internal/db"
	"time"
)

func Migrate(conn db.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	migrationList := []migrations.Migration{
		migrations.MigrationUserAdmin,
		migrations.MigrationEmployee,
		migrations.MigrationCompanies,
	}

	for _, migration := range migrationList {
		_ = migration(ctx, conn)
	}

	return nil
}
