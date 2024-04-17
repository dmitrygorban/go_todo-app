package server

import (
	"log"
	"net/http"

	"github.com/dmitrygorban/go_todo-app/handlers"
	"github.com/dmitrygorban/go_todo-app/middlewares"
	"github.com/dmitrygorban/go_todo-app/storage"
)

type Server struct {
	httpServer  *http.Server
	taskHandler *handlers.TaskHandler
}

func NewServer(port string, taskHandler *handlers.TaskHandler) *Server {
	server := &Server{
		taskHandler: taskHandler,
		httpServer: &http.Server{
			Addr: port,
		},
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/api/signin", handlers.SignInHandler)
	http.HandleFunc("/api/nextdate", handlers.GetNextDate)
	http.HandleFunc("/api/task", middlewares.Auth(s.taskHandler.HandleTaskRequests))
	http.HandleFunc("/api/tasks", middlewares.Auth(s.taskHandler.GetTasks))
	http.HandleFunc("/api/task/done", middlewares.Auth(s.taskHandler.DoneTask))
}

func (s *Server) Start(store storage.TaskStore) error {

	log.Printf("Server starting on port %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
