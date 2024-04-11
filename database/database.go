package database

import (
	"database/sql"
	"os"
	"path/filepath"
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
