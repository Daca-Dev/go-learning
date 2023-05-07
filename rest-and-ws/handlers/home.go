package handlers

import (
	"encoding/json"
	"net/http"

	"daca-code.com/go/rest-ws/server"
)

type HomeResponse struct {
	Message string `json:"message"` // el json:"" es la forma que le decimos a go que esta serializando
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to platzi GO",
			Status:  true,
		})
	}
}
