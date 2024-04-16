package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *TaskHandler) edit(w http.ResponseWriter, r *http.Request) {
	task, err := parseRequest(r)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		handleError(w, "title is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.ParseInt(task.Id, 10, 0)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.Store.GetTaskById(int(idInt))
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	newDate, err := calculateDateToSave(task)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.Date = newDate

	err = h.Store.Update(task)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = emptyOkResponse(w)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
	}
}

func emptyOkResponse(w http.ResponseWriter) error {

	response := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
