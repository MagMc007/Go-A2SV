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
type UserRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserService {
	return &UserRepository{
		collection: collection,
	}
}

func (u *UserRepository) GetByUsername(username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := u.collection.FindOne(
		ctx, 
		bson.M{"username": username
	}).Decode(&user)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepository) Register(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

