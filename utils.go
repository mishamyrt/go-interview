package main

import (
	"errors"
	"fmt"
	"regexp"
)

// findTaskByID ищет задачу по ID в срезе задач.
func findTaskByID(tasks []Task, id string) (*Task, error) {
	find := func(t Task) bool {
		return t.ID == id
	}
	for _, task := range tasks {
		if find(task) {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

// recursiveTaskFormatter выполняет рекурсивное форматирование для списка задач, добавляя префиксы к каждому описанию.
func recursiveTaskFormatter(tasks []Task, prefix string) []string {
	if len(tasks) == 0 {
		return []string{}
	}

	formattedTask := fmt.Sprintf("%s %s: %s", prefix, tasks[0].ID, tasks[0].Description)
	return append([]string{formattedTask}, recursiveTaskFormatter(tasks[1:], prefix)...)
}

// sanitizeTaskDescription очищает описание задачи от запрещённых символов.
func sanitizeTaskDescription(description *string) error {
	re := regexp.MustCompile(`[!@#$%^&*()_+={}|[\]\\:;"'<>,.?/]`)
	*description = re.ReplaceAllString(*description, "")
	return nil
}

// taskStatusToString возвращает строковое представление статуса задачи.
func taskStatusToString(completed bool) string {
	if completed {
		return "Completed"
	}
	return "Pending"
}
