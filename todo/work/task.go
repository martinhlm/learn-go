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

type TaskManager struct {
	task *Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Save(task *Task) error {
	m.task = task
	return nil
}

func (m *TaskManager) All() []*Task {
	return []*Task{m.task}
}
