package server

import (
	"log"
	"net/http"

	"github.com/dmitrygorban/go_todo-app/handlers"
	"github.com/dmitrygorban/go_todo-app/middlewares"
	"github.com/dmitrygorban/go_todo-app/storage"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: port,
		},
	}
}

func (s *Server) Start(store storage.TaskStore) error {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	taskHandler := handlers.NewTaskHandler(store)

	http.HandleFunc("/api/signin", handlers.SignInHandler)
	http.HandleFunc("/api/nextdate", handlers.GetNextDate)
	http.HandleFunc("/api/task", middlewares.Auth(taskHandler.HandleTaskRequests))
	http.HandleFunc("/api/tasks", middlewares.Auth(taskHandler.GetTasks))
	http.HandleFunc("/api/task/done", middlewares.Auth(taskHandler.DoneTask))
	log.Printf("Server starting on port %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
