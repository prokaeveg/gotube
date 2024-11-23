package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func Dsn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		),
	))
}
