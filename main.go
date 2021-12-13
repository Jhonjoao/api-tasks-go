package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

type Task struct {
	ID            int
	Name          string `json:"name"`
	Description   string `json:"description"`
	OwnerTask     string `json:"email"`
	Status        string `json:"status"` // maybe change this to bool
	TimesFinished int
}

var Tasks []Task

var PENDING_STATUS string = "pending"
var FINISHED_STATUS string = "finished"

func removeTask(s []Task, index int) []Task {
	return append(s[:index], s[index+1:]...)
}

func setTaskId(task Task) int {
	if len(Tasks) == 0 {
		return 1
	} else {
		return Tasks[len(Tasks)-1].ID + 1
	}
}

func updateTaskStatus(task Task, status string) Task {
	if status == PENDING_STATUS && task.Status == FINISHED_STATUS {
		task.TimesFinished = task.TimesFinished + 1
	}
	task.Status = status
	return task
}

func validStatus(status string) bool {
	if status == PENDING_STATUS || status == FINISHED_STATUS {
		return true
	}
	return false
}

func resJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func listTasks(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var selectedType = params["type"]
	if validStatus(selectedType) {
		var searchTasks []Task
		for _, item := range Tasks {
			if item.Status == selectedType {
				searchTasks = append(searchTasks, item)
			}
		}
		resJSON(res, http.StatusOK, searchTasks)
		return
	}
	resJSON(res, http.StatusNoContent, "")
}

func createTask(res http.ResponseWriter, req *http.Request) {
	var task Task
	_ = json.NewDecoder(req.Body).Decode(&task)

	// set ID
	task.ID = setTaskId(task)
	task.TimesFinished = 0

	// set default status
	if task.Status == "" {
		task.Status = "pending"
	}

	// validate status
	if !validStatus(task.Status) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid Status"))
		return
	}

	Tasks = append(Tasks, task)
	resJSON(res, http.StatusCreated, task)
}

func updateTask(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var task Task
	_ = json.NewDecoder(req.Body).Decode(&task)

	if !validStatus(task.Status) {
		resJSON(res, http.StatusBadRequest, "Invalid Status")
		return
	}

	taskId, _ := strconv.Atoi(params["taskID"])

	var foundTask, index = searchTaskById(taskId)

	if foundTask.ID == 0 {
		resJSON(res, http.StatusBadRequest, "Not found task with ID "+params["taskID"])
		return
	}

	if foundTask.TimesFinished == 3 && task.Status == PENDING_STATUS {
		resJSON(res, http.StatusBadRequest, "It is not possible to move the task to 'pending' more than 3 times")
		return
	}

	Tasks[index] = updateTaskStatus(foundTask, task.Status)

	resJSON(res, http.StatusOK, Tasks[index])
}

func searchTaskById(taskID int) (Task, int) {
	var t Task
	var position int

	for index, item := range Tasks {
		if item.ID == taskID {
			t = item
			position = index
			break
		}
	}

	return t, position
}

func deleteTask(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range Tasks {
		intVar, _ := strconv.Atoi(params["taskID"])
		if item.ID == intVar {
			Tasks = removeTask(Tasks, index)
			resJSON(res, http.StatusNoContent, "")
		}
	}

}

//

func (a *App) Run(addr string) {
	fmt.Println("Server at " + addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/task/{type}", listTasks).Methods(http.MethodGet)
	a.Router.HandleFunc("/task", createTask).Methods(http.MethodPost)
	a.Router.HandleFunc("/task/{taskID:[0-9]+}", updateTask).Methods(http.MethodPatch)
	a.Router.HandleFunc("/task/{taskID:[0-9]+}", deleteTask).Methods(http.MethodDelete)
}
