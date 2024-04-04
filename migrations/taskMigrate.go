package migrations

import (
	"log"

	"github.com/dmitrygorban/go_todo-app/database"
)

func TaskMigrate(store database.TaskStore) {
	createTable := `CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    date CHAR(8) NOT NULL DEFAULT "",
    title TEXT NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
  )`

	createIndex := `CREATE INDEX task_date ON scheduler(date)`

	_, err := store.Db.Exec(createTable)

	if err != nil {
		log.Fatalf("Error creating scheduler table: %v", err)
	}

	_, err = store.Db.Exec(createIndex)

	if err != nil {
		log.Fatalf("Error creating index in scheduler table: %v", err)
	}
}
