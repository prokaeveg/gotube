package auth

import (
	"context"
	"encoding/json"
	"gotube/api"
	"net/http"
)

func HandleAuthorization(repo UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = UserAuthRequest{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			api.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		if request.Login == "" || request.Password == "" {
			api.RespondError(w, http.StatusBadRequest, "Login and/or password are required")
			return
		}

		response, err := AuthUser(context.Background(), repo, request)

		if err != nil {
			api.RespondError(w, http.StatusBadRequest, "Invalid credentials")
			return
		}

		api.RespondSuccess(w, response)
	}
}
