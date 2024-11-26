package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gotube/auth"
	"gotube/db"
	"gotube/handlers"
	"gotube/mock"
	"net/http"
	"strconv"
)

func (server *Server) MountTestHandlers() {
	router := server.Router

	router.Route("/test", func(router chi.Router) {
		router.Route("/user/{id:[0-9]+}", func(router chi.Router) {
			router.Get("/token", func(writer http.ResponseWriter, r *http.Request) {
				id, _ := strconv.Atoi(chi.URLParam(r, "id"))
				token, refreshToken := auth.CreateTokenForUser(id)

				writer.Header().Set("Content-Type", "application/json")
				_, _ = writer.Write([]byte(fmt.Sprintf("{\"token\":\"%s\",\"refresh_token\":\"%s\"}", token, refreshToken)))
			})
		})

		router.Post("/auth", func(writer http.ResponseWriter, r *http.Request) {
			fakeResponse, _ := mock.AuthUser("test@home.su", "FloorFloor123")

			writer.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(writer).Encode(fakeResponse)
		})

		router.Get("/dsn", handlers.Dsn)
		router.Get("/greet", handlers.Greet)
		router.Get("/sql", func(w http.ResponseWriter, r *http.Request) {
			db.ShowTables(server.DBRepo.DB, w) //@todo вместо DB передавать репозиторий
		})
	})
}
