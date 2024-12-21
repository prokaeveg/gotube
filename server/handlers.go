package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gotube/admin"
	"gotube/api"
	"gotube/auth"
	"gotube/file"
	"net/http"
)

func (server *Server) MountHandlers() {
	router := server.Router
	router.Use(middleware.CleanPath)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		response := map[string]string{
			"message": "Hello, World!",
		}
		api.RespondSuccess(writer, response)
	})

	router.Route("/file", func(r chi.Router) {
		r.Get("/form", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("<!DOCTYPE html>\n<html>\n  <head>\n    <title>File Upload</title>\n  </head>\n  <body>\n    <h1>File Upload</h1>\n    <form action=\"/upload\" method=\"post\" enctype=\"multipart/form-data\">\n      <input type=\"file\" name=\"file\" />\n      <br /><br />\n      <input type=\"submit\" value=\"Upload\" />\n    </form>\n  </body>\n</html>"))
		})
		r.Post("/upload", file.HandleUploadedFile())
	})

	// Admin routes
	router.Route("/admin", func(router chi.Router) {
		router.Route("/users", func(router chi.Router) {
			router.Get("/", admin.UserListHandler(server.DBRepo.DB)) //@todo вместо DB передавать репозиторий
		})
		router.Post("/create", admin.CreateUserHandler(server.DBRepo.DB))
	})

	// Auth
	server.Router.Post("/authorize", auth.HandleAuthorization(&server.DBRepo))
}
