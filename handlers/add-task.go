package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dmitrygorban/go_todo-app/models"
	"github.com/dmitrygorban/go_todo-app/scheduler"
)

func (h *TaskHandler) add(w http.ResponseWriter, r *http.Request) {

	newTask, err := parseRequest(r)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newTask.Title == "" {
		handleError(w, "title is required", http.StatusBadRequest)
		return
	}

	dateToSave, err := calculateDateToSave(newTask)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Store.Add(models.Task{
		Title:   newTask.Title,
		Date:    dateToSave,
		Repeat:  newTask.Repeat,
		Comment: newTask.Comment,
	})

	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, id)
}
func parseRequest(r *http.Request) (models.Task, error) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	return task, err
}

func calculateDateToSave(task models.Task) (string, error) {
	now := time.Now().Truncate(24 * time.Hour)
	var dateTime time.Time
	var err error

	if task.Date != "" {
		dateTime, err = time.Parse(scheduler.DATE_FORMAT, task.Date)
		if err != nil {
			return "", err
		}
	} else {
		dateTime = now
	}

	if dateTime.Before(now) && task.Repeat != "" {
		nextDate, err := scheduler.NextDate(now, dateTime.Format(scheduler.DATE_FORMAT), task.Repeat)
		if err != nil {
			return "", err
		}
		return nextDate, nil
	} else if dateTime.Before(now) && task.Repeat == "" {
		return now.Format(scheduler.DATE_FORMAT), nil
	}

	return dateTime.Format(scheduler.DATE_FORMAT), nil
}

func sendResponse(w http.ResponseWriter, id int) {
	response := models.Task{
		Id: strconv.Itoa(id),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
