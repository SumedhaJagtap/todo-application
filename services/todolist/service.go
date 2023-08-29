package todolist

import (
	"fmt"

	"github.com/todoapplication/db"
	"github.com/todoapplication/services/task"
)

func listTasks() (tasks []task.Task, err error) {

	rows, err := db.GetDB().Query("SELECT * from tasks")
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.CreatedAt, &t.ModifiedAt); err != nil {
			return tasks, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return tasks, fmt.Errorf("FetchAllAdmins %v", err)
	}
	return tasks, nil
}
