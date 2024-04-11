package database

import (
	"database/sql"

	"github.com/dmitrygorban/go_todo-app/models"
)

func (s *TaskStore) UpdateTask(task models.Task) error {

	_, err := s.Db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.Id),
	)
	if err != nil {
		return err
	}

	return nil
}
