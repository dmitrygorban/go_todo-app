package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dmitrygorban/go_todo-app/models"
)

func handleError(w http.ResponseWriter, errorMessage string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(code)

	errorResponse := models.ErrorResponse{Error: errorMessage}

	err := json.NewEncoder(w).Encode(errorResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
