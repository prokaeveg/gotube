package server

import (
	"github.com/go-chi/chi/v5"
	"gotube/admin"
	"gotube/db"
	"gotube/handlers"
	"net/http"
)

func (server *Server) MountHandlers() {
	server.Router.Route("/admin", func(router chi.Router) {
		router.Get("/users", admin.UserListHandler(server.DB))
	})
	server.Router.Get("/dsn", handlers.Dsn)
	server.Router.Get("/greet", handlers.Greet)
	server.Router.Get("/sql", func(w http.ResponseWriter, r *http.Request) {
		db.ShowTables(server.DB, w)
	})
}
