package work

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("Learn Go")
	if task.Title != "Learn Go" {
		t.Errorf("expected learn Go, got %v", task.Title)
	}
}

func TestNewTaskWithEmptyTitle(t *testing.T) {
	task := NewTask("")
	if task != nil {
		t.Errorf("expected nil, got %#v", task)
	}
}
