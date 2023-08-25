package todorest

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/todoapplication/todo"
)

type todoListDB struct {
	db *sql.DB
}

func InitTodoList() todo.TodoList {

	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createTableQuery := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		name TEXT,
		status BOOLEAN
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	return &todoListDB{db: db}
}

func (tasklist *todoListDB) AddTask(name string) {
	query := "INSERT INTO tasks (name,status) values(?,?)"
	_, err := tasklist.db.Exec(query, name, 0)
	if err != nil {
		log.Fatal(err)
	}
}
func (tasklist *todoListDB) ListTasks() {
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
func (tasklist *todoListDB) DeleteTask(index int) {
	deleteQuery := "DELETE from tasks where  ID=?"
	_, err := tasklist.db.Exec(deleteQuery, index)
	if err != nil {
		log.Fatal(err)
	}
}

func (tasklist *todoListDB) MarkAsDone(index int) {
	updateQuery := "UPDATE tasks set status=1 where  ID=?"
	_, err := tasklist.db.Exec(updateQuery, index)
	if err != nil {
		log.Fatal(err)
	}
}

func (tasklist *todoListDB) Run() {
	// createTableQuery := `CREATE TABLE IF NOT EXISTS tasks (
	// 	id INTEGER PRIMARY KEY,
	// 	name TEXT,
	// 	status BOOLEAN
	// );`

	// _, err := tasklist.db.Exec(createTableQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	task := flag.String("task", "", "Add task")
	list := flag.Bool("list", false, "List all task")
	done := flag.Int("done", 0, "Mark task as Done")
	delete := flag.Int("delete", 0, "Delete task")
	flag.Parse()
	switch {
	case *task != "":
		fmt.Println("Add Task")
		tasklist.AddTask(*task)
		fmt.Println("Task added.")
	case *list:
		tasklist.ListTasks()
	case *done != 0:
		tasklist.MarkAsDone(*done)
	case *delete != 0:
		tasklist.DeleteTask(*delete)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}

}
