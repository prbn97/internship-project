package task

import (
	"api/main.go/types"
	"api/main.go/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Store struct {
	filename string
	tasks    map[string]types.Task
}

func NewStore(filename string) (*Store, error) {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	store := &Store{
		filename: filename,
		tasks:    make(map[string]types.Task),
	}
	err := store.loadTasks()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) loadTasks() error {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &s.tasks)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) saveTasks() error {
	data, err := json.Marshal(s.tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateTask(payload types.TaskPayLoad) error {
	id, err := utils.GenerateID(20)
	if err != nil {
		return err
	}

	task := types.Task{
		ID:          id,
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   false,
	}

	s.tasks[id] = task
	err = s.saveTasks()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListTasks() ([]*types.Task, error) {
	tasks := make([]*types.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (s *Store) GetTaskByID(id string) (*types.Task, error) {
	task, exists := s.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}
	return &task, nil
}

func (s *Store) UpdateTask(task types.Task) error {
	_, exists := s.tasks[task.ID]
	if !exists {
		return fmt.Errorf("task not found")
	}

	s.tasks[task.ID] = task
	return s.saveTasks()
}

func (s *Store) DeleteTask(id string) (types.Task, error) {
	task, exists := s.tasks[id]
	if !exists {
		return task, fmt.Errorf("task not found")
	}

	delete(s.tasks, id)
	err := s.saveTasks()
	if err != nil {
		return task, err
	}
	return task, nil
}
