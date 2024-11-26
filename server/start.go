package server

import (
	"log"
	"net/http"
	"os"
)

func (server *Server) Start() {
	port := os.Getenv("APP_PORT")
	log.Printf("Server running on %s", port)
	if err := http.ListenAndServe(":"+port, server.Router); err != nil {

		log.Fatalf("Failed to start server: %v", err)
	}
}
