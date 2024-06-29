package types

import "time"

// Users Payloads
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Users Object
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

// Users Interface
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// Tasks Payloads
type TaskPayLoad struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

// Tasks Object
type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Tasks Interface
type TaskStore interface {
	CreateTask(TaskPayLoad) error
	ListTasks() ([]*Task, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(Task) error
	DeleteTask(id string) (Task, error)
}
