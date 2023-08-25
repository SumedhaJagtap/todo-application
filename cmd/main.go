package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/todoapplication/todo"
	"github.com/todoapplication/todo_rest/server"
)

type TodoRunner todo.TodoList

func main() {

	// tasklist := todorest.InitTodoList()

	// todo.Run(tasklist)
	server.StartServer()

}
