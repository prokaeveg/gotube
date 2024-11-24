package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Router *chi.Mux
	DB     *pgxpool.Pool
}

func CreateServer(db *pgxpool.Pool) *Server {
	server := &Server{
		Router: chi.NewRouter(),
		DB:     db,
	}

	return server
}
