package main

import (
	"github.com/joho/godotenv"
	"gotube/db"
	"gotube/migrations"
	"gotube/server"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pool := db.InitializeDb()
	srv := server.CreateServer(pool)
	srv.MountHandlers()
	migrations.RunMigrations(srv.DB)

	srv.Start()
}
