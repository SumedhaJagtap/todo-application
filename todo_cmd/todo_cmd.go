package todo_cmd

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/todoapplication/todo"
)

type TodoListCMD struct {
	filename string
	tasks    []todo.Task //slice
}

func InitTodoList(filename string) todo.TodoList {
	return &TodoListCMD{filename: filename}
}

func (tasklist *TodoListCMD) AddTask(name string) {
	t := todo.Task{Name: name, Status: false}
	tasklist.tasks = append(tasklist.tasks, t)
}

func (tasklist *TodoListCMD) ListTasks() {
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
func (tasklist *TodoListCMD) DeleteTask(index int) {
	if index >= 1 && index <= len(tasklist.tasks) {
		fmt.Println("Remove ", index)
		tasklist.tasks = append(tasklist.tasks[:index-1], tasklist.tasks[index:]...)
	} else {
		fmt.Println("Index not found")
	}
}

func (tasklist *TodoListCMD) MarkAsDone(index int) {
	if index >= 1 && index <= len(tasklist.tasks) {
		tasklist.tasks[index-1].Status = true
	} else {
		fmt.Println("Index not found")
	}
}

func (tasklist *TodoListCMD) Run() {
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
		tasklist.AddTask(*task)
		fmt.Println("Task added.")
		json, err := json.Marshal(tasklist.tasks)
		fmt.Println(tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(tasklist.filename, json, 0644)
	case *list:
		tasklist.ListTasks()

	case *done != 0:
		tasklist.MarkAsDone(*done)
		json, err := json.Marshal(tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(tasklist.filename, json, 0644)
	case *delete != 0:

		tasklist.DeleteTask(*delete)
		json, err := json.Marshal(tasklist.tasks)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(tasklist.filename, json, 0644)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}
}
