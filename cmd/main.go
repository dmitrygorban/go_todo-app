package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dmitrygorban/go_todo-app/migrations"
	"github.com/dmitrygorban/go_todo-app/server"
	"github.com/dmitrygorban/go_todo-app/storage"
)

const DEFAULT_PORT = ":7540"
const DEFAULT_DB_PATH = "./scheduler.db"

func main() {
	portToListen := DEFAULT_PORT
	dbPath := DEFAULT_DB_PATH

	envPort := os.Getenv("TODO_PORT")
	if envPort != "" {
		portToListen = fmt.Sprintf(":%s", envPort)
	}

	envPathToDb := os.Getenv("DB_PATH")
	if envPathToDb != "" {
		dbPath = fmt.Sprintf(":%s", envPathToDb)
	}

	store := storage.NewTaskStore(dbPath)
	defer store.Db.Close()

	if store.NeedMigration(dbPath) {
		migrations.TaskMigrate(store)
	}

	httpServer := server.NewServer(portToListen)
	err := httpServer.Start(store)

	if err != nil {
		log.Fatal(err)
	}
}
