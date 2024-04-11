package database

import (
	"database/sql"

	"github.com/dmitrygorban/go_todo-app/models"
)

func (s *TaskStore) Add(t models.Task) (int, error) {
	res, err := s.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
