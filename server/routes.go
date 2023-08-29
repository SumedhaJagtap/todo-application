package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/todoapplication/services/task"
	"github.com/todoapplication/services/todolist"
)

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", defaultHandler).Methods(http.MethodGet)
	router.HandleFunc("/task", task.AddTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks", todolist.ListTasks).Methods(http.MethodGet)

	router.HandleFunc("/task/{taskid}", task.DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/task/{taskid}/complete", task.DoneTask).Methods(http.MethodPost)

	return
}
