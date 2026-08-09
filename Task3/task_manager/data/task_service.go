package data

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"task_manager/models"
)

type TaskServices interface {
	AddTask(task models.Task) models.Task
	GetAllTasks() []models.Task
	GetTaskDetails(id int) (models.Task, error)
	UpdateTask(id int, task models.Task) (models.Task, error)
	DeleteTask(id int) error
}

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		collection: collection,
	}
}

func (t *TaskService) AddTask(task models.Task) models.Task {
	t.tasks = append(t.tasks, task)

	return task
}

func (t *TaskService) GetAllTasks() []models.Task {
	return t.tasks
}

func (t *TaskService) GetTaskDetails(id int) (models.Task, error) {
	for _, v := range t.tasks {
		if id == v.ID {
			return v, nil
		}
	}

	return models.Task{}, errors.New("task with this ID does not exist")
}

func (t *TaskService) UpdateTask(id int, task models.Task) (models.Task, error) {
	for i, v := range t.tasks {
		if id == v.ID {
			t.tasks[i] = task
			return task, nil
		}
	}

	return models.Task{}, errors.New("task with this ID does not exist")
}

func (t *TaskService) DeleteTask(id int) error {
	for i, v := range t.tasks {
		if id == v.ID {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("task with this ID does not exist")
}