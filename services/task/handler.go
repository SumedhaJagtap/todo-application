package task

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	server_utility "github.com/todoapplication/server/utility"
	utility "github.com/todoapplication/utility"
)

func AddTask(w http.ResponseWriter, req *http.Request) {
	var task = &Task{}
	err := json.NewDecoder(req.Body).Decode(&task)
	switch {
	case err == io.EOF:
		http.Error(w, server_utility.ErrRequestBodyEmpty.Error(), http.StatusBadRequest)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.Status = 0
	task.CreatedAt = sql.NullString{
		String: strconv.FormatInt(utility.GetEpochTime(), 16),
		Valid:  true,
	}
	taskId, err := addTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server_utility.Success(http.StatusCreated, server_utility.Response{Message: fmt.Sprintf("User #%d updated Successfully.", taskId)}, w)
}

func DeleteTask(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)    // Extract URL variables
	taskid := vars["taskid"] // Get the value of the "variable" variable
	if taskid == "" {
		server_utility.Error(http.StatusBadRequest, server_utility.Response{Message: server_utility.ErrTaskEmptyID.Error()}, w)
	}
	id, err := strconv.ParseInt(taskid, 10, 64)
	if err != nil {
		server_utility.Error(http.StatusBadRequest, server_utility.Response{Message: err.Error()}, w)
		return
	}
	err = deleteTask(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server_utility.Success(http.StatusAccepted, server_utility.Response{Message: fmt.Sprintf("Task #%d deleted Successfully.", id)}, w)
}

func DoneTask(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)    // Extract URL variables
	taskid := vars["taskid"] // Get the value of the "variable" variable
	if taskid == "" {
		server_utility.Error(http.StatusBadRequest, server_utility.Response{Message: server_utility.ErrTaskEmptyID.Error()}, w)
	}
	id, err := strconv.ParseInt(taskid, 10, 64)
	if err != nil {
		server_utility.Error(http.StatusBadRequest, server_utility.Response{Message: err.Error()}, w)
		return
	}
	err = markAsDone(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server_utility.Success(http.StatusAccepted, server_utility.Response{Message: fmt.Sprintf("Task #%d updated Successfully.", id)}, w)
}
