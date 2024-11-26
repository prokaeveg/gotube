package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gotube/admin"
	"gotube/auth"
	"net/http"
)

func (server *Server) MountHandlers() {
	router := server.Router
	router.Use(middleware.CleanPath)

	router.Route("/admin", func(router chi.Router) {
		router.Route("/users", func(router chi.Router) {
			router.Get("/", admin.UserListHandler(server.DBRepo.DB)) //@todo вместо DB передавать репозиторий
		})
		router.Post("/create", func(writer http.ResponseWriter, r *http.Request) {

		})
	})

	server.Router.Post("/authorize", auth.HandleAuthorization(&server.DBRepo))
}
