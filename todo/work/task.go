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
