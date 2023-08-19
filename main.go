package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type TodoList interface {
	run()
	addTask(name string)
	listTasks()
	deleteTask(index int)
	markAsDone(index int)
}

type TodoListCMD struct {
	tasks []Task
}

type TodoListDB struct {
	db    *sql.DB
	tasks []Task
}

type TodoListfile struct {
	filename string
	tasks    []Task
}

type Task struct {
	Name   string
	Status bool
}

func (tasklist *TodoListCMD) addTask(name string) {
	t := Task{Name: name, Status: false}
	tasklist.tasks = append(tasklist.tasks, t)
}

func (tasklist *TodoListCMD) listTasks() {
	if len(tasklist.tasks) == 0 {
		fmt.Println("No tasks in the list.")
	} else {
		fmt.Println("-----Task List-----")
		fmt.Println("Task\tStatus")
		for i, task := range tasklist.tasks {
			fmt.Printf("%d. %s - %v\n", i+1, task.Name, task.Status)
		}
	}
}

func (tasklist *TodoListCMD) deleteTask(index int) {
	if index >= 1 && index <= len(tasklist.tasks) {
		fmt.Println("Remove ", index)
		tasklist.tasks = append(tasklist.tasks[:index-1], tasklist.tasks[index:]...)
	} else {
		fmt.Println("Index not found")
	}
}

func (tasklist *TodoListCMD) markAsDone(index int) {
	if index >= 1 && index <= len(tasklist.tasks) {
		tasklist.tasks[index-1].Status = true
	} else {
		fmt.Println("Index not found")
	}
}

func (tasklist *TodoListCMD) run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n-----TO-DO LIST-----")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Quit")
		fmt.Println("Enter your choice : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Add Task")
			scanner.Scan()
			task := scanner.Text()
			tasklist.addTask(task)
			fmt.Println("Task added.")
		case "2":
			tasklist.listTasks()
		case "3":
			fmt.Println("Enter the task index to mark as done: ")
			scanner.Scan()
			indexStr := scanner.Text()
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				fmt.Println("Invalid input please enter valid index")
				continue
			}
			tasklist.markAsDone(index)
		case "4":
			fmt.Println("Enter the task index to delete: ")
			scanner.Scan()
			indexStr := scanner.Text()
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				fmt.Println("Invalid input please enter valid index")
				continue
			}
			tasklist.deleteTask(index)
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (tasklist *TodoListfile) addTask(name string) {
	t := Task{Name: name, Status: false}
	tasklist.tasks = append(tasklist.tasks, t)
}
func (tasklist *TodoListfile) listTasks() {
	if len(tasklist.tasks) == 0 {
		fmt.Println("No tasks in the list.")
	} else {
		fmt.Println("-----Task List-----")
		fmt.Println("Task\tStatus")
		for i, task := range tasklist.tasks {
			fmt.Printf("%d. %s - %v\n", i+1, task.Name, task.Status)
		}
	}
}
func (tasklist *TodoListfile) deleteTask(index int) {
	if index >= 1 && index <= len(tasklist.tasks) {
		fmt.Println("Remove ", index)
		tasklist.tasks = append(tasklist.tasks[:index-1], tasklist.tasks[index:]...)
	} else {
		fmt.Println("Index not found")
	}
}

func (tasklist *TodoListfile) markAsDone(index int) {
	if index >= 1 && index <= len(tasklist.tasks) {
		tasklist.tasks[index-1].Status = true
	} else {
		fmt.Println("Index not found")
	}
}
func (tasklist *TodoListfile) run() {
	task := flag.String("task", "", "Add task")
	list := flag.Bool("list", false, "List all task")
	done := flag.Int("done", 0, "Mark task as Done")
	delete := flag.Int("delete", 0, "Delete task")
	file, err := ioutil.ReadFile(tasklist.filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File not exists")
			os.Exit(0)
		}
		fmt.Println(err)
		os.Exit(0)
	}
	if len(file) != 0 {
		err = json.Unmarshal(file, &tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	flag.Parse()

	switch {
	case *task != "":
		tasklist.addTask(*task)
		fmt.Println("Task added.")
		json, err := json.Marshal(tasklist.tasks)
		fmt.Println(tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(tasklist.filename, json, 0644)
	case *list:
		tasklist.listTasks()

	case *done != 0:
		tasklist.markAsDone(*done)
		json, err := json.Marshal(tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(tasklist.filename, json, 0644)
	case *delete != 0:

		tasklist.deleteTask(*delete)
		json, err := json.Marshal(tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(tasklist.filename, json, 0644)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}

}

func (tasklist *TodoListDB) addTask(name string) {
	query := "INSERT INTO tasks (name,status) values(?,?)"
	_, err := tasklist.db.Exec(query, name, 0)
	if err != nil {
		log.Fatal(err)
	}
}
func (tasklist *TodoListDB) listTasks() {
	rows, err := tasklist.db.Query("SELECT name,status from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var task Task
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
func (tasklist *TodoListDB) deleteTask(index int) {
	deleteQuery := "DELETE from tasks where  ID=?"
	_, err := tasklist.db.Exec(deleteQuery, index)
	if err != nil {
		log.Fatal(err)
	}
}

func (tasklist *TodoListDB) markAsDone(index int) {
	updateQuery := "UPDATE tasks set status=1 where  ID=?"
	_, err := tasklist.db.Exec(updateQuery, index)
	if err != nil {
		log.Fatal(err)
	}
}
func (tasklist *TodoListDB) run() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		name TEXT,
		status BOOLEAN
	);`

	_, err := tasklist.db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("\n-----TO-DO LIST-----")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Quit")
		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var taskTitle string
			fmt.Println("Add Task")
			fmt.Scanln(&taskTitle)
			tasklist.addTask(taskTitle)
			fmt.Println("Task added.")
		case "2":
			tasklist.listTasks()
		case "3":
			fmt.Println("Enter the task index to mark as done: ")
			var index int
			fmt.Scanln(&index)
			tasklist.markAsDone(index)
		case "4":
			fmt.Println("Enter the task index to delete: ")
			var index int
			fmt.Scanln(&index)
			tasklist.deleteTask(index)
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
func runTodoApplication(t TodoList) {
	t.run()
}

func main() {
	//<--- array datastore and user inputs as cmds --->///
	// todoList := TodoListCMD{}
	// todoList.run()

	//<--- array datastore and cmdline args --->///
	todoList := TodoListfile{filename: "tasks.txt"}
	todoList.run()

	//<--- DB as sqlite3 files --->///
	// db, err := sql.Open("sqlite3", "todo.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// 	todoList := TodoListDB{db: db}
	// todoList.run()
}
