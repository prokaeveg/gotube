package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

type Server struct {
	Router *chi.Mux
	DBRepo DBRepository
}

type DBRepository struct {
	DB *pgxpool.Pool
}

func (r *DBRepository) FindUserIdByCredentials(ctx context.Context, username string, password string) (int, error) {
	var userID int
	//@todo сделать хэширование пароля
	err := r.DB.QueryRow(ctx, "SELECT id FROM users WHERE username = $1 AND password_hash = $2", username, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func CreateServer(db *pgxpool.Pool) *Server {
	DBRepo := DBRepository{
		DB: db,
	}
	server := &Server{
		Router: chi.NewRouter(),
		DBRepo: DBRepo,
	}

	return server
}
