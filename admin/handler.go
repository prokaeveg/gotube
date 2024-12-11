package admin

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"gotube/api"
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

func CreateUserHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser = CreateUserRequest{}

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			api.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		var validation = validator.New()

		if err := validation.Struct(newUser); err != nil {
			//TODO: extract normal errors from validation
			api.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		err := CreateUser(db, newUser)

		if err != nil {
			api.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		api.RespondSuccess(w, newUser)
	}
}
