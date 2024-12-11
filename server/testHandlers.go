package server

import (
	"github.com/go-chi/chi/v5"
	"gotube/api"
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

				response := map[string]string{
					"token":         token,
					"refresh_token": refreshToken,
				}
				api.RespondSuccess(writer, response)
			})
		})

		router.Post("/auth", func(writer http.ResponseWriter, r *http.Request) {
			fakeResponse, _ := mock.AuthUser("test@home.su", "FloorFloor123")

			api.RespondSuccess(writer, fakeResponse)
		})

		router.Get("/dsn", handlers.Dsn)
		router.Get("/greet", handlers.Greet)
		router.Get("/sql", func(w http.ResponseWriter, r *http.Request) {
			db.ShowTables(server.DBRepo.DB, w) //@todo вместо DB передавать репозиторий
		})
	})
}
