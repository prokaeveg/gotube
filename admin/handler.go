package admin

import (
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func UserListHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := FetchAllUsers(db)
		if err != nil {
			http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
