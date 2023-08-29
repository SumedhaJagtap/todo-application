package task

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/todoapplication/db"
	server_utility "github.com/todoapplication/server/utility"
	utility "github.com/todoapplication/utility"
)

func addTask(ctx context.Context, task *Task) (taskId int64, err error) {
	if task.ID != 0 {
		return 0, fmt.Errorf("TaskID should be empty")
	}
	if task.Name == "" {
		return 0, fmt.Errorf("Task name should be non empty")
	}

	result, err := db.GetDB().Exec(fmt.Sprintf("INSERT INTO tasks(name,description,status,createdAt) VALUES('%s','%s','%v','%v')", task.Name, task.Description, task.Status, task.CreatedAt))
	if err != nil {
		return 0, fmt.Errorf("AddTask: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddTask: %v", err)
	}
	return id, nil
}

func GetTaskById(ctx context.Context, taskId int64) (*Task, error) {
	var task Task
	if taskId == 0 {
		return nil, fmt.Errorf("TaskID should be non-empty")
	}
	row := db.GetDB().QueryRow(fmt.Sprintf("SELECT * from tasks where ID=%d", taskId))
	if err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.CreatedAt, &task.ModifiedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, server_utility.NoTaskFound
		}
		return nil, err
	}
	return &task, nil
}
func deleteTask(ctx context.Context, taskId int64) error {
	if taskId == 0 {
		return fmt.Errorf("TaskID should be non-empty")
	}
	_, err := GetTaskById(ctx, taskId)
	if err != nil {
		return fmt.Errorf("deleteTask: %v", err)
	}
	_, err = db.GetDB().Exec(fmt.Sprintf("DELETE FROM tasks where ID = %d", taskId))
	if err != nil {
		return fmt.Errorf("deleteTask: %v", err)
	}
	return nil
}

func markAsDone(ctx context.Context, taskId int64) error {
	if taskId == 0 {
		return fmt.Errorf("TaskID should be non-empty")
	}
	_, err := GetTaskById(ctx, taskId)
	if err != nil {
		return fmt.Errorf("MarkAsDone: %v", err)
	}
	modifiedAt := sql.NullString{
		String: strconv.FormatInt(utility.GetEpochTime(), 16),
		Valid:  true,
	}
	_, err = db.GetDB().Exec(fmt.Sprintf("UPDATE tasks set Status=1,ModifiedAt='%v' where ID = %d", modifiedAt, taskId))
	if err != nil {
		return fmt.Errorf("MarkAsDone: %v", err)
	}
	return nil
}
