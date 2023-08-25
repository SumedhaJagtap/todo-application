package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/todoapplication/todo"
	"github.com/todoapplication/todo_cmd"
	todorest "github.com/todoapplication/todo_rest"
)

type TodoRunner todo.TodoList

func main() {

	var todoRunner TodoRunner
	tasklist := todo_cmd.InitTodoList("tasks.txt")
	todoRunner = tasklist

	todorest.InitTodoList()

	todoRunner.Run()

	// db, err := sql.Open("sqlite3", "todo.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// tasklist = todorest.InitTodoList(db)
	// todo.Run(tasklist)
}
