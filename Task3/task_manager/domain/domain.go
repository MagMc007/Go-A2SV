package	domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// structs
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Duedate     time.Time          `bson:"dueDate" json:"dueDate"`
	Status      string             `bson:"status" json:"status"`
}


type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"`
	Role     string             `bson:"role" json:"role"`
}

// interfaces
type TaskRepository interface {
	AddTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskDetails(id primitive.ObjectID) (Task, error)
	UpdateTask(id primitive.ObjectID, task Task) (Task, error)
	DeleteTask(id primitive.ObjectID) error
}

type TaskUsecase interface {
	AddTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskDetails(id primitive.ObjectID) (Task, error)
	UpdateTask(id primitive.ObjectID, task Task) (Task, error)
	DeleteTask(id primitive.ObjectID) error
}

// usecase
type UserRepository interface {
	Register(user User) (User, error)
	GetByUsername(username string) (User, error)
}


type UserUsecase interface {
    Register(user User) (User, error)
    Login(username string, password string) (string, error)
}