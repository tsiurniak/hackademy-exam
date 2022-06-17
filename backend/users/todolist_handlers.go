package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddListParams struct {
	Name string `json:"name"`
}

func AddList(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	params := AddListParams{}
	err := readParams(r, &params)

	if err != nil {
		HandleError(err, rw)
		return
	}

	id := u.NewToDoList(params.Name)

	list, err := u.Get(id)

	if err != nil {
		HandleError(err, rw)
		return
	}

	marshalled, err := list.MarshalJSON()

	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusCreated, string(marshalled))
}

type UpdateListParams AddListParams

func UpdateList(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	params := UpdateListParams{}
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

	list.name = params.Name

	marshalled, err := list.MarshalJSON()

	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusCreated, string(marshalled))
}

func DeleteList(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["list_id"])

	if err != nil {
		HandleError(err, rw)
		return
	}

	_, err = u.Delete(uint(id))

	if err != nil {
		HandleError(err, rw)
		return
	}

	writeResponse(rw, http.StatusNoContent, "")
}

func GetLists(rw http.ResponseWriter, r *http.Request, u *UserStorage) {
	keys := make([]uint, 0)
	for key := range u.lists {
		keys = append(keys, key)
	}

	lists := make([]map[string]string, 0)
	for key := range keys {
		list, _ := u.Get(uint(key))
		object := make(map[string]string)

		object["id"] = fmt.Sprint(list.id)
		object["name"] = list.name

		lists = append(lists, object)
	}

	fmt.Println(lists)

	marshalled, err := json.Marshal(lists)
	if err != nil {
		HandleError(err, rw)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	writeResponse(rw, http.StatusOK, string(marshalled))
}
