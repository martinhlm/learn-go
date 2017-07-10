package work

import "testing"

func newTaskOrFatal(t *testing.T, title string) *Task {
	task, err := NewTask(title)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	return task
}

func TestNewTask(t *testing.T) {
	task := newTaskOrFatal(t, "Learn Go")
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

func TestSaveTask(t *testing.T) {
	task := newTaskOrFatal(t, "Learn Go")

	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSaveTaskAndRetrive(t *testing.T) {
	task := newTaskOrFatal(t, "Learn Go")

	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	all := m.All()
	if len(all) != 1 {
		t.Fatalf("expected one task, got %v", len(all))
	}
	if all[0].Title != task.Title {
		t.Errorf("expected title %q, got %q", task.Title, all[0].Title)
	}
}
