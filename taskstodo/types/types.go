package types

// Users

// Tasks
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
