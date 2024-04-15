package migrations

import (
	"log"

	"github.com/dmitrygorban/go_todo-app/storage"
)

func TaskMigrate(store storage.TaskStore) {
	createTable := `CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    date CHAR(8) NOT NULL, 
    title TEXT NOT NULL,
    comment TEXT DEFAULT NULL, 
    repeat VARCHAR(128) DEFAULT NULL 
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
