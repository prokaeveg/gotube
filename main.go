package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"gotube/db"
	"gotube/handlers"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Router *chi.Mux
	DB     *pgxpool.Pool
}

func CreateServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}

	return server
}

func (server *Server) InitDB() {
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

	server.DB = pool
}

func (server *Server) MountHandlers() {
	server.Router.Get("/dsn", handlers.Dsn)
	server.Router.Get("/greet", handlers.Greet)
	server.Router.Get("/sql", func(w http.ResponseWriter, r *http.Request) {
		db.ShowTables(server.DB, w)
	})
}

func runMigrations(pool *pgxpool.Pool) {
	migrationDir := "./migrations"

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	sqlConnection := stdlib.OpenDBFromPool(pool)

	if err := goose.Up(sqlConnection, migrationDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := CreateServer()
	server.InitDB()
	server.MountHandlers()

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", server.Router)

	if err != nil {
		panic(err)
	}

	runMigrations(server.DB)
}
