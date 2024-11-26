package main

import (
	"github.com/joho/godotenv"
	"gotube/db"
	"gotube/migrations"
	"gotube/server"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pool := db.InitializeDb()
	srv := server.CreateServer(pool)
	srv.MountHandlers()

	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "test" {
		srv.MountTestHandlers()
	}

	migrations.RunMigrations(srv.DBRepo.DB)

	srv.Start()
}
