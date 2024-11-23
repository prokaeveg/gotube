package handlers

import "net/http"

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func Sql(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}
