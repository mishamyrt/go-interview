package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/task/add", addTaskHandler)
	http.HandleFunc("/task/remove", removeTaskHandler)
	http.HandleFunc("/task/edit", editTaskHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/export", exportTasksHandler)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	log.Println("Starting server on :8080")
	server.ListenAndServe()
}
