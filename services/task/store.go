package task

import (
	"cmd/main.go/types"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func generateID(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid length")
	}
	numBytes := length / 2
	if length%2 != 0 {
		numBytes++
	}
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	id := hex.EncodeToString(randomBytes)

	if len(id) > length {
		id = id[:length]
	} else if len(id) < length {
		id += strings.Repeat("0", length-len(id))
	}

	return id, nil
}

// func TestGenerateID(t *testing.T) {
// 	id, err := GenerateID(20)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 20, len(id))
// 	_, err = GenerateID(-25)
// 	assert.Error(t, err)
// }

func (s *Store) CreateTask(payload types.TaskPayLoad) error {
	id, err := generateID(20)
	if err != nil {
		return err
	}

	task := types.Task{
		ID:          id,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      false,
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
