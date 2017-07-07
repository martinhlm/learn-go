package work

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("Learn Go")
	if task.Title != "Learn Go" {
		t.Errorf("expected learn Go, got %v", task.Title)
	}
}
