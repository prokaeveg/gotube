package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gotube/db"
	"gotube/handlers"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Router *chi.Mux
	DB     *sql.DB
}

func CreateServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}

	return server
}

func (server *Server) InitDB() {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	server.DB = db
}

func (server *Server) MountHandlers() {
	server.Router.Get("/greet", handlers.Greet)
	server.Router.Get("/sql", func(w http.ResponseWriter, r *http.Request) {
		db.ShowTables(server.DB, w)
	})
}

func main() {
	server := CreateServer()
	server.MountHandlers()

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", server.Router)

	if err != nil {
		panic(err)
	}
}
