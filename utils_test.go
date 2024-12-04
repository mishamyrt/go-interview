package main

import (
	"testing"
)

func TestFindTaskByID(t *testing.T) {
	tasks := []Task{
		{ID: "1", Description: "Test Task 1"},
		{ID: "2", Description: "Test Task 2"},
	}

	_, err := findTaskByID(tasks, "1")
	if err != nil {
		t.Errorf("Expected to find task with ID '1', got '%v'", err)
	}

	_, err = findTaskByID(tasks, "3")
	if err != nil {
		t.Errorf("Expected to find task with ID '3', got error: %v", err)
	}
}

func TestRecursiveTaskFormatter(t *testing.T) {
	tasks := []Task{
		{ID: "1", Description: "First Task", Completed: true},
		{ID: "2", Description: "Second Task", Completed: false},
	}

	formattedTasks := recursiveTaskFormatter(tasks)
	if len(formattedTasks) != 2 {
		t.Errorf("Expected 2 formatted tasks, got %d", len(formattedTasks))
	}
	if formattedTasks[0].Completed != true {
		t.Error("Expected 'Task 1' to be completed")
	}
}

func TestSanitizeTaskDescription(t *testing.T) {
	description := "!@#Test$%^&*Task"
	expected := "Test Task"

	err := sanitizeTaskDescription(&description)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if description != expected {
		t.Errorf("Expected description to be '%s', got '%s'", expected, description)
	}
	err = sanitizeTaskDescription(nil)
	if err == nil {
		t.Errorf("Expected error when description is nil")
	}
}
