package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/todoapplication/todo"
	"github.com/todoapplication/todo_cmd"
)

type TodoRunner todo.TodoList

func main() {

	tasklist := todo_cmd.InitTodoList("tasks.txt")

	todo.Run(tasklist)

}
