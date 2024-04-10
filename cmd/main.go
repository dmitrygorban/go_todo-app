package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dmitrygorban/go_todo-app/database"
	"github.com/dmitrygorban/go_todo-app/handlers"
	"github.com/dmitrygorban/go_todo-app/migrations"
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
	fs := http.FileServer(http.Dir("./web"))

	http.Handle("/", fs)

	db := database.NewSqlliteDatabase(dbPath)
	defer db.Db.Close()

	storage := database.NewTaskStore(db.Db)
	taskHandler := handlers.NewTaskHandler(storage)

	install := database.DoesDbInstallRequired(dbPath)
	if install {
		migrations.TaskMigrate(storage)
	}

	http.HandleFunc("/api/nextdate", handlers.GetNextDate)
	http.HandleFunc("/api/task", taskHandler.HandleTaskRequests)

	log.Printf("Server starting on port %s", portToListen)
	err := http.ListenAndServe(portToListen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
