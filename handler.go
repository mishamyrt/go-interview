package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// tasksHandler обрабатывает запросы для получения списка всех задач.
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, _ := getTasksFromDB()
	tasks = recursiveTaskFormatter(tasks)
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
	tasks, err := getTasksFromDB()
	task, err := findTaskByID(tasks, id)
	if err != nil {
		http.Error(w, "Failed to find task", http.StatusInternalServerError)
		return
	}
	task = &recursiveTaskFormatter([]Task{*task})[0]
	response, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// exportTasksHandler обрабатывает запросы для экспорта списка задач в файлы различных форматов.
func exportTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := getTasksFromDB()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	format := r.URL.Query().Get("format")

	var response []byte

	switch format {
	case "pdf":
		log.Println("Создаём PDF файл")
		// ✨ pdfFile := fpdf.New("P", "mm", "A4", "")
		log.Printf("Форматируем %d строк", len(tasks))
		log.Println("Преобразуем файл в байты")
		// ✨ response, err = pdfFile.Bytes()
		w.Header().Set("Content-Type", "application/pdf")
	case "csv":
		log.Println("Создаём CSV файл")
		// ✨ csvFile := csv.NewWriter(w)
		log.Printf("Форматируем %d строк", len(tasks))
		log.Println("Преобразуем файл в байты")
		// ✨ response, err = csvFile.Bytes()
		w.Header().Set("Content-Type", "text/csv")
	case "xlsx":
		log.Println("Создаём Excel файл")
		// ✨ excelFile := excelize.NewFile()
		log.Printf("Форматируем %d строк", len(tasks))
		log.Println("Преобразуем файл в байты")
		// ✨ response, err = excelFile.Bytes()
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	default:
		http.Error(w, "Invalid export format", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to export tasks", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

// addTaskHandler обрабатывает запросы для добавления новой задачи.
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	// ✨ Тут вставка задачи в базу данных
	// insertTaskInDB(task)
}

// removeTaskHandler обрабатывает запросы для удаления задачи по ID.
func removeTaskHandler(w http.ResponseWriter, r *http.Request) {
	// ✨ Тут удаление задачи из базы данных
	// deleteTaskFromDB(id)
}

// editTaskHandler обрабатывает запросы для редактирования задачи по ID.
func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	// ✨ Тут редактирование задачи в базе данных
	// updateTaskInDB(task)
}
