package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/dmitrygorban/go_todo-app/models"
)

type TaskStore struct {
	Db *sql.DB
}

func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{Db: db}
}

func DoesDbInstallRequired(pathToFile string) bool {
	appPath, err := os.Getwd()

	if err != nil {
		return false
	}

	dbFile := filepath.Join(appPath, pathToFile)
	_, err = os.Stat(dbFile)

	if os.IsNotExist(err) {
		return true
	}

	return false
}

func (s TaskStore) Add(t models.Task) (int, error) {
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
