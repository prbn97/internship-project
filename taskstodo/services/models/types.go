package types

type TaskPayLoad struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskStore interface {
	CreateTask(TaskPayLoad) error
	ListTasks() ([]*Task, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(Task) error
	DeleteTask(id string) (Task, error)
}
