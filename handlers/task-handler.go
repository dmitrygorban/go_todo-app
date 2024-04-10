package handlers

import "github.com/dmitrygorban/go_todo-app/database"

type TaskHandler struct {
	Store database.TaskStore
}

func NewTaskHandler(store database.TaskStore) *TaskHandler {
	return &TaskHandler{Store: store}
}
