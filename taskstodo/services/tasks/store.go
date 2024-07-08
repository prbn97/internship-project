package tasks

import (
	"database/sql"
	"fmt"

	types "github.com/prbn97/internship-project/services/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTask(payload types.TaskPayLoad, userID int) error {

	_, err := s.db.Exec("INSERT INTO tasks (userID, title, description) VALUES (?, ?, ?)", userID, payload.Title, payload.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ListTasks(userID int) ([]*types.Task, error) {

	rows, err := s.db.Query("SELECT * FROM tasks WHERE userID = ?", userID)
	if err != nil {
		return nil, err
	}

	tasks := make([]*types.Task, 0)
	for rows.Next() {
		task, err := scanRowsIntoTask(rows)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) GetTaskByID(userID int, id int) (*types.Task, error) {

	row := s.db.QueryRow("SELECT * FROM tasks WHERE userID = ? AND id = ?", userID, id)

	task := new(types.Task)
	err := row.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("task %d not found", id)

	}

	return task, nil
}

func (s *Store) UpdateTask(task types.Task) error {

	_, err := s.db.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE userID = ? AND id = ?", task.Title, task.Description, task.Status, task.UserID, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteTask(userID int, id int) error {

	result, err := s.db.Exec("DELETE FROM tasks WHERE userID = ? AND id = ?", userID, id)
	if err != nil {
		return fmt.Errorf("could not delete task %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not verify rows affected for task %d: %v", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("task %d not found for user %d", id, userID)
	}

	return nil
}

func scanRowsIntoTask(rows *sql.Rows) (*types.Task, error) {
	task := new(types.Task)

	err := rows.Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}
