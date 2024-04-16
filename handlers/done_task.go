package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dmitrygorban/go_todo-app/scheduler"
)

func (h *TaskHandler) DoneTask(w http.ResponseWriter, r *http.Request) {
	queryId := r.FormValue("id")

	idInt, err := strconv.ParseInt(queryId, 10, 0)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.Store.GetTaskById(int(idInt))
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		h.delete(w, r)
		return
	}

	nextRepeatDate, err := scheduler.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task.Date = nextRepeatDate

	err = h.Store.Update(task)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = emptyOkResponse(w)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
