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

// recursiveTaskFormatter выполняет рекурсивное форматирование для списка задач, добавляя индикатор статуса к каждому описанию.
func recursiveTaskFormatter(tasks []Task) []Task {
	if len(tasks) == 0 {
		return []Task{}
	}

	var mark string
	if tasks[0].Completed {
		mark = "✅"
	} else {
		mark = "🛑"
	}
	description := fmt.Sprintf("%s: %s", mark, tasks[0].Description)

	return append([]Task{
		{
			ID:          tasks[0].ID,
			Description: description,
			Completed:   tasks[0].Completed,
		},
	}, recursiveTaskFormatter(tasks[1:])...)
}

// sanitizeTaskDescription очищает описание задачи от запрещённых символов.
func sanitizeTaskDescription(description *string) error {
	re := regexp.MustCompile(`[!@#$%^&*()_+={}|[\]\\:;"'<>,.?/]`)
	*description = re.ReplaceAllString(*description, "")
	return nil
}
