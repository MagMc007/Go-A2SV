package data

import (
	"errors"

	"task_manager/models"
)

// prepare services for CRUD ops
// prepare an interface

type TaskServices interface {
	AddTask(task models.Task) models.Task
	GetAllTasks() []models.Task

	GetTaskDetails(id int) (models.Task, error)
	UpdateTask(id int, task model.Task) (models.Task, error)
	DeleteTask(id int) error
}

func NewTaskService() *TaskServices {
	return &TaskServices{
		tasks: make([]models.Task, 0),
	}
}

// extend each method in the interface and implement
func (t *[]models.Task) AddTask(task models.Task) models.Task {
	*t = append(*t, task)

	return task
}

func (t *[]models.Task) GetAllTasks() []models.Task{
	return *t
}

func (t *[]models.Task) GetTaskDetails(id int) (models.Task, error) {
	for _, v := range *t {
		if id == v.ID {
			return (v, nil)
		}
	}

	return (models.Task{}, errors.New("Task with this ID does not exist"))
}


func (t *[]models.Task) UpdateTask(id int, task models.Task) (models.Task, error) {
	for i, v := range *t {
		if id == v.ID {
			(*t)[i] = task	
		}
	}

	return (models.Task{}, errors.New("Task with this ID does not exist"))	
}

func (t *[]models.Task) DeleteTask(id int) error {
	for i, v := range *t {
		if id == v.ID {
			(*t) = append((*t)[:i], (*t)[i+1:]...)
		}
	}

	return (errors.New("Task with this ID does not exist"))	
}