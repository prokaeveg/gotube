package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func InitializeDb() *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	config, err := pgxpool.ParseConfig(dsn)

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//if err := pool.Ping(context.Background()); err != nil {
	//	log.Fatalf("Failed to ping database: %v", err)
	//}

	return pool
}
