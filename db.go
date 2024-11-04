package main

import (
	"database/sql"
	"errors"
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
	defer rows.Close()

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
	_, err = db.Exec("INSERT INTO tasks (id, description, completed) VALUES ($1, $2, $3)", task.ID, task.Description, task.Completed)
	if err != nil {
		return err
	}
	return nil
}

// deleteTaskFromDB удаляет задачу из базы данных по ID.
func deleteTaskFromDB(id string) error {
	result, err := db.Exec(fmt.Sprintf("DELETE FROM tasks WHERE id = '%s'", id))
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted, task not found")
	}

	return nil
}

// updateTaskStatusInDB обновляет статус выполнения задачи в базе данных.
func updateTaskStatusInDB(id string, completed bool) error {
	_, err := db.Exec("UPDATE tasks SET completed = $1 WHERE id = $2", completed, id)
	return err
}
