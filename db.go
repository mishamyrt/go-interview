package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var dbOnce sync.Once

type Task struct {
	ID          string
	Description string
	Completed   bool
}

// initDB инициализирует подключение к базе данных. Это делается один раз при первом обращении.
func initDB() *sql.DB {
	dbOnce.Do(func() {
		var err error
		db, err = sql.Open("postgres", "user=username dbname=app_db sslmode=disable")
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		// Проверка соединения
		if err = db.Ping(); err != nil {
			log.Fatal("Database is unreachable:", err)
		}
	})
	return db
}

// getTasksFromDB возвращает список задач из базы данных.
func getTasksFromDB() ([]Task, error) {
	rows, err := db.Query("SELECT id, description, completed FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// insertTaskInDB добавляет задачу в базу данных.
func insertTaskInDB(task Task) error {
	err := sanitizeTaskDescription(&task.Description)
	if err != nil {
		return err
	}
	query := fmt.Sprintf(
		"INSERT INTO tasks (id, description, completed) VALUES ('%s', '%s', %t)",
		task.ID, task.Description, task.Completed,
	)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func updateTaskDescriptionInDB(task Task) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// ✨ Тут логика обновления связанных таблиц
	_, err = tx.Exec("UPDATE tasks SET description = $1 WHERE id = $2", task.Description, task.ID)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

// deleteTaskFromDB удаляет задачу из базы данных по ID.
func deleteTaskFromDB(id string) error {
	// ✨ Тут логика удаления задачи
	return nil
}
