package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddTaskParams struct {
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func AddTask(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	params := AddTaskParams{}
	err := readParams(r, &params)

	if err != nil {
		HandleError(err, rw)
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["list_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	list, err := u.Get(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	task_id, err := list.NewTask(params.TaskName, params.Description)

	if err != nil {
		HandleError(err, rw)
		return
	}

	task, err := list.Get(task_id)

	if err != nil {
		HandleError(err, rw)
		return
	}

	marshalled, err := task.MarshalJSON()

	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusCreated, string(marshalled))
}

type UpdateTaskParams struct {
	AddTaskParams
	UpdateName        bool `json:"update_name"`
	UpdateDescription bool `json:"update_description"`
	UpdateStatus      bool `json:"update_status"`
}

func (t *Task) update(params UpdateTaskParams) {
	if params.UpdateName {
		t.name = params.TaskName
	}
	if params.UpdateDescription {
		t.description = params.Description
	}
	if params.UpdateStatus {
		t.status = params.Status
	}
}

func UpdateTask(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	params := UpdateTaskParams{}
	err := readParams(r, &params)

	if err != nil {
		HandleError(err, rw)
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["list_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	list, err := u.Get(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	id, err = strconv.Atoi(vars["task_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	task, err := list.Get(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	task.update(params)

	err = list.Update(task.id, task)

	if err != nil {
		HandleError(err, rw)
		return
	}

	marshalled, err := task.MarshalJSON()

	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusCreated, string(marshalled))
}

func DeleteTask(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["list_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	list, err := u.Get(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	id, err = strconv.Atoi(vars["task_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	_, err = list.Delete(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusNoContent, "")
}

func GetTasks(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["list_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	list, err := u.Get(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	keys := make([]uint, 0)
	for key := range list.tasks {
		keys = append(keys, key)
	}

	tasks := make([]map[string]string, 0)
	for _, key := range keys {
		task, err := list.Get(uint(key))

		if err != nil {
			HandleError(err, rw)
			return
		}

		object := make(map[string]string)

		object["id"] = fmt.Sprint(task.id)
		object["task_name"] = task.name
		object["description"] = task.description
		object["open"] = task.status

		tasks = append(tasks, object)
	}

	fmt.Println(tasks)

	marshalled, err := json.Marshal(tasks)
	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusOK, string(marshalled))
}
