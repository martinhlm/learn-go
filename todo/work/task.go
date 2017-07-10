package work

import "errors"

type Task struct {
	Title string
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, errors.New("empty title")
	}
	return &Task{title}, nil
}

type TaskManager struct{}

func NewTaskManager() *TaskManager {
	return nil
}

func (m *TaskManager) Save(task *Task) error {
	return nil
}
