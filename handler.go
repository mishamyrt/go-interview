package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// tasksHandler обрабатывает запросы для получения списка всех задач.
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, _ := getTasksFromDB()
	response, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// taskHandler обрабатывает запросы для получения задачи по ID.
func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	tasks, _ := getTasksFromDB()
	task, err := findTaskByID(tasks, id)
	if err != nil {
		http.Error(w, "Failed to find task", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// taskStatusHandler обрабатывает запросы для получения статуса выполнения задачи по ID.
func taskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	tasks, _ := getTasksFromDB()
	task, err := findTaskByID(tasks, id)
	if err != nil {
		http.Error(w, "Failed to find task", http.StatusInternalServerError)
		return
	}
	status := taskStatusToString(task.Completed)
	w.Write([]byte(status))
}

// addTaskHandler обрабатывает запросы для добавления новой задачи.
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var task Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	err = insertTaskInDB(task)
	if err != nil {
		http.Error(w, "Failed to add task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task added successfully"))
}

// removeTaskHandler обрабатывает запросы для удаления задачи по ID.
func removeTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	err := deleteTaskFromDB(idParam)
	if err != nil {
		http.Error(w, "Failed to remove task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task removed successfully"))
}

func prefixTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, _ := getTasksFromDB()
	prefix := r.URL.Query().Get("id")
	taskNames := recursiveTaskFormatter(tasks, prefix)
	response, err := json.Marshal(taskNames)
	if err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
