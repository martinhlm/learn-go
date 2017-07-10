package work

type Task struct {
	Title string
}

func NewTask(title string) *Task {
	if title == "" {
		return nil
	}
	return &Task{title}
}
