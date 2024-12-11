package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func RespondSuccess(w http.ResponseWriter, data interface{}) {
	response := Response{
		Status: "success",
		Error:  "",
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RespondError(w http.ResponseWriter, statusCode int, error string) {
	response := Response{
		Status: "error",
		Error:  error,
		Data:   nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
