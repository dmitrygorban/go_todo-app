package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *TaskHandler) getById(w http.ResponseWriter, r *http.Request) {
	queryId := r.FormValue("id")

	id, err := strconv.ParseInt(queryId, 10, 0)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.Store.GetTaskById(int(id))
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
	}
}
