package main

import (
	"log"
	"solid-task-management/internal/di"
)

func main() {
	server, err := di.InitializeServer()
	if err != nil {
		log.Fatalf("could not initialize server: %v", err)
	}

	server.Run()
}
