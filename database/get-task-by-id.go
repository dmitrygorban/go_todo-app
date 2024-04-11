package database

import (
	"database/sql"

	"github.com/dmitrygorban/go_todo-app/models"
)

func (s *TaskStore) GetTaskById(id int) (models.Task, error) {
	task := models.Task{}

	row := s.Db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id;",
		sql.Named("id", id))

	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return task, err
	}

	return task, nil
}
