package task

import "database/sql"

type Task struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      int            `json:"status"`
	CreatedAt   sql.NullString `json:"created_at"`
	ModifiedAt  sql.NullString `json:"modified_at"`
}

type UserService interface {
	addTask(task *Task) (taskId int64, err error)
	deleteTask(taskId int64) error
	GetTaskById(taskId int64) (*Task, error)
	markAsDone(taskId int64) error
}
