package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const DEFAULT_PORT = ":7540"

func main() {
	portToListen := DEFAULT_PORT

	envPort := os.Getenv("TODO_PORT")

	fs := http.FileServer(http.Dir("./web"))

	http.Handle("/", fs)

	if envPort != "" {
		portToListen = fmt.Sprintf(":%s", envPort)
	}

	log.Printf("Server starting on port %s", portToListen)

	err := http.ListenAndServe(portToListen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
