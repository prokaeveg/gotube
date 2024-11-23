package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Router *chi.Mux
}

func CreateServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}

	return server
}

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func (server *Server) MountHandlers() {
	server.Router.Get("/greet", Greet)
}

func main() {
	server := CreateServer()
	server.MountHandlers()

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", server.Router)

	if err != nil {
		panic(err)
	}
}
