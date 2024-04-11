package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dmitrygorban/go_todo-app/models"
)

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	tasks, err := h.Store.Get(search)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	response := models.TasksResponse{
		Tasks: tasks,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
	}
}
