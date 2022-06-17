package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type Task struct {
	id          uint
	name        string
	description string
	status      string
}

func (t Task) MarshalJSON() ([]byte, error) {
	object := make(map[string]string)
	object["id"] = fmt.Sprint(t.id)
	object["task_name"] = t.name
	object["description"] = t.description
	object["status"] = t.status

	marshalled, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	return marshalled, nil
}

type TaskFactory struct {
	currId uint
}

func NewTaskFactory() TaskFactory {
	return TaskFactory{
		currId: 1,
	}
}

func (tf *TaskFactory) NewTask(name, description string) Task {
	defer func() { tf.currId += 1 }()

	return Task{
		id:          tf.currId,
		name:        name,
		description: description,
		status:      "open",
	}
}

type ToDoList struct {
	id   uint
	name string
	// List unique identifier
	tasks   map[uint]Task // Map taskId -> Task
	lock    sync.RWMutex
	factory TaskFactory
}

func (tdl *ToDoList) MarshalJSON() ([]byte, error) {
	object := make(map[string]string)
	object["id"] = fmt.Sprint(tdl.id)
	object["name"] = tdl.name

	marshalled, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	return marshalled, nil
}

func (tl *ToDoList) NewTask(name, description string) (uint, error) {
	newTask := tl.factory.NewTask(name, description)
	err := tl.Add(newTask.id, newTask)

	if err != nil {
		return 0, err
	}
	return newTask.id, nil
}

type ToDoListFactory TaskFactory

func (tlf *ToDoListFactory) NewToDoList(name string) (*ToDoList, uint) {
	defer func() { tlf.currId += 1 }()

	return &ToDoList{
		id:      tlf.currId,
		name:    name,
		tasks:   make(map[uint]Task),
		factory: TaskFactory{currId: 1},
	}, tlf.currId
}

func (l *ToDoList) Add(key uint, task Task) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.tasks[key].id != 0 {
		return errors.New("Key '" + fmt.Sprint(key) + "' already exists")
	}

	l.tasks[key] = task
	return nil
}

func (l *ToDoList) Update(key uint, task Task) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.tasks[key].id == 0 {
		return errors.New("Key '" + fmt.Sprint(key) + "' doesn't exist")
	}

	l.tasks[key] = task
	return nil
}

func (l *ToDoList) Get(key uint) (task Task, err error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	task, exists := l.tasks[key]
	if exists {
		return task, nil
	}
	return (Task{}), errors.New("Key '" + fmt.Sprint(key) + "' doesn't exist")
}

func (l *ToDoList) Delete(key uint) (task Task, err error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	task, exists := l.tasks[key]
	if exists {
		delete(l.tasks, key)
		return task, nil
	}
	return (Task{}), errors.New("Key '" + fmt.Sprint(key) + "' doesn't exist")
}
