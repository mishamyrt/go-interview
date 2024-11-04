package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/task/status", taskStatusHandler)
	http.HandleFunc("/task/add", addTaskHandler)
	http.HandleFunc("/task/remove", removeTaskHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/prefix", prefixTasksHandler)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}
