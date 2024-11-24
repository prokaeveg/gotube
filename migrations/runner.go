package migrations

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

func RunMigrations(pool *pgxpool.Pool) {
	migrationDir := "./migrations"

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	sqlConnection := stdlib.OpenDBFromPool(pool)

	if err := goose.Up(sqlConnection, migrationDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
