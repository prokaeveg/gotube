package auth

import (
	"context"
	"encoding/json"
	"net/http"
)

func HandleAuthorization(repo UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = UserAuthRequest{}

		json.NewDecoder(r.Body).Decode(&request)
		response, err := AuthUser(context.Background(), repo, request)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
