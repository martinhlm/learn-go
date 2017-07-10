package work

import "testing"

func TestNewTask(t *testing.T) {
	task, err := NewTask("Learn Go")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if task.Title != "Learn Go" {
		t.Errorf("expected learn Go, got %v", task.Title)
	}
}

func TestNewTaskWithEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Errorf("expected 'empty title' error, got %v", err)
	}
}
