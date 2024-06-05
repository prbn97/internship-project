# April day 22 - *What is a API?*
Last Friday I received the task of developing a Rest API.

We received a recommendation from a friend to use OpenAPI to plan the entire API,
Therefore, my focus was to better understand the concept and history of the API, RestFul API.

OpenAPI, in short, plans an entire API, in addition to the documentation. This produces a professional application. However, the focus is not on the principle of an API.

# June day 04 -  *Milestone 5: Documentation and API Refinement*

On day 3 I spoke to my tutor, I showed him an architecture of an API separated into layers to be more modularized.
The study was very interesting. I found a guy called [Tiago](https://www.youtube.com/watch?v=OVwUldzmVOg&t=353s), Portuguese with great content in Go.
However, I will try this structure at another time... I saved this [playlist](https://www.youtube.com/watch?v=7VLmLOiQ3ck&list=PLYEESps429vqQ98y_zjyERFQR1cyvBNzA&index=4) for the future.

However, I made improvements to the file structure and API code.
## API Structure
    /todo-api/
            + /cmd/
                - /main.go

            + /db/
                - /tasksLists.json

            + /docs/
                - /guideLine.md
                - /dayBookDevelopment.md
                - /documentation.md

            + /services/
                    + /task/
                        - /routes.go
                        - /store.go
            + /types/                        
                - /types.go

            + /utils/                        
                - /utils.go

            - README.md
            - go.mod
            - Makefile


## API logic

An API where it is possible to create tasks (tasks)
with the following information.
```go
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
```

### in folder /services/
We will have the ability to create tasks, it is divided into:

> **routes.go** to handle HTTP requests and handle errors.
In this file we will have the **Handler type** that will receive the interface. **TaskStore** and thus communicate with store.go.

```go
type Handler struct {
	store types.TaskStore
}

type TaskStore interface {
	CreateTask(TaskPayLoad) error
	ListTasks() ([]*Task, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(Task) error
	DeleteTask(id string) (Task, error)
}
```

> **store.go** to handle storing the created tasks in a json file.
In this file we have the type **Store** accesses the json file where it is saved

```go
type Store struct {
	filename string
	tasks    map[string]types.Task
}
```

### in folder /types/
> you will find the structures and interfaces that services use.
```go
package types

type TaskPayLoad struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskStore interface {
	CreateTask(TaskPayLoad) error
	ListTasks() ([]*Task, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(Task) error
	DeleteTask(id string) (Task, error)
}
```

### in folder /utils/
> you will find the functions that serve all API services.
```go
var Validate = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {}

func WriteError(w http.ResponseWriter, status int, err error) {}

func ParseJSON(r *http.Request, v any) error {}
```
