package handlers

import "github.com/dmitrygorban/go_todo-app/storage"

type TaskHandler struct {
	Store storage.TaskStore
}

func NewTaskHandler(store storage.TaskStore) *TaskHandler {
	return &TaskHandler{Store: store}
}
