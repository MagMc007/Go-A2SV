package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
    collection *mongo.Collection
}


func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{
		collection: collection,
	}
}

type UserServices interface {
	Register(user models.User) (models.User, error)
}

func (u *UserService) Register(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check for duplicate names
	err := u.collection.FindOne(
		ctx, 
		bson.M{"username": user.Username},
	).Err()

	if err == nil {
		return models.User{}, errors.New("username already exists")
	}

	if err != mongo.ErrNoDocuments {
		return models.User{}, err
	}

	// hash pwd
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	// insert the user
	_, err = u.collection.InsertOne(ctx, user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}