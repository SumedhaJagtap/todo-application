package task

import (
	"fmt"
	"log"

	"github.com/todoapplication/todo"
)

type TodoListApplication interface {
	AddTask(name string)
	ListTasks()
	DeleteTask(index int)
	MarkAsDone(index int)
}

func (tasklist *TodoList) AddTask(name string) {
	query := "INSERT INTO tasks (name,status) values(?,?)"
	_, err := tasklist.db.Exec(query, name, 0)
	if err != nil {
		log.Fatal(err)
	}
}
func (tasklist *TodoList) ListTasks() {
	rows, err := tasklist.db.Query("SELECT name,status from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tasks []todo.Task
	for rows.Next() {
		var task todo.Task
		if err := rows.Scan(&task.Name, &task.Status); err != nil {
			log.Fatal(err)

		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
	} else {
		fmt.Println("-----Task List-----")
		fmt.Println("Task\tStatus")
		for i, task := range tasks {
			fmt.Printf("%d. %s - %v\n", i+1, task.Name, task.Status)
		}
	}
}
func (tasklist *TodoList) DeleteTask(index int) {
	deleteQuery := "DELETE from tasks where  ID=?"
	_, err := tasklist.db.Exec(deleteQuery, index)
	if err != nil {
		log.Fatal(err)
	}
}

func (tasklist *TodoList) MarkAsDone(index int) {
	updateQuery := "UPDATE tasks set status=1 where  ID=?"
	_, err := tasklist.db.Exec(updateQuery, index)
	if err != nil {
		log.Fatal(err)
	}
}
