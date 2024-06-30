package tasks

import (
	"database/sql"

	"github.com/prbn97/internship-project/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTask(payload types.TaskPayLoad) error {

	return nil
}

func (s *Store) ListTasks() ([]*types.Task, error) {

	return nil, nil
}

func (s *Store) GetTaskByID(id string) (*types.Task, error) {

	return nil, nil
}

func (s *Store) UpdateTask(task types.Task) error {

	return nil
}

func (s *Store) DeleteTask(id string) (types.Task, error) {

	var mock types.Task
	return mock, nil
}
