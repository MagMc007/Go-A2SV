package repositories

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"task_manager/domain"
)


/*
	this layer maintains interaction with DB, only CRUD, nothing more,
	Also no business logic, validation ...
*/


type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		collection: collection,
	}
}


func (t *TaskRepository) AddTask(task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.collection.InsertOne(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}


func (t *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := t.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskRepository) GetTaskDetails(id primitive.ObjectID) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task domain.Task

	err := t.collection.FindOne(
		ctx,
		bson.M{"_id": id},
	).Decode(&task)

	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (t *TaskRepository) UpdateTask(id primitive.ObjectID, task domain.Task) (domain.Task, error) {
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
		return domain.Task{}, err
	}

	if result.MatchedCount == 0 {
		return domain.Task{}, errors.New("task not found")
	}

	task.ID = id

	return task, nil
}

func (t *TaskRepository) DeleteTask(id primitive.ObjectID) error {
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