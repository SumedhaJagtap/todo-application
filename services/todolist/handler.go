package todolist

import (
	"net/http"

	server_utility "github.com/todoapplication/server/utility"
)

func ListTasks(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	tasks, err := listTasks(ctx)
	if err != nil {
		server_utility.Error(http.StatusInternalServerError, server_utility.Response{Message: err.Error()}, w)
	}
	server_utility.Success(http.StatusAccepted, server_utility.Response{Message: "tasks", Details: tasks}, w)

}
