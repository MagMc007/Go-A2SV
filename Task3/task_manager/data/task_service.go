package data

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"task_manager/models"
)

type TaskServices interface {
	AddTask(task models.Task) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskDetails(id primitive.ObjectID) (models.Task, error)
	UpdateTask(id primitive.ObjectID, task models.Task) (models.Task, error)
	DeleteTask(id primitive.ObjectID) error
}

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		collection: collection,
	}
}

func (t *TaskService) AddTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.collection.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (t *TaskService) GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := t.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) GetTaskDetails(id primitive.ObjectID) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task

	err := t.collection.FindOne(
		ctx,
		bson.M{"_id": id},
	).Decode(&task)

	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (t *TaskService) UpdateTask(id primitive.ObjectID, task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"dueDate":     task.Duedate,
			"status":      task.Status,
		},
	}

	result, err := t.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		return models.Task{}, err
	}

	if result.MatchedCount == 0 {
		return models.Task{}, errors.New("task not found")
	}

	task.ID = id

	return task, nil
}

func (t *TaskService) DeleteTask(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := t.collection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}