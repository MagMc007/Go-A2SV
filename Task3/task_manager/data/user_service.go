package data

import (
	"context"
	"errors"
	"fmt"
	"os"
	"task_manager/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	Login(username string, password string) (string, error)
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

func (u *UserService) Login(username string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	// Find the user
	err := u.collection.FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid username or password")
		}

		return "", err
	}

	// fmt.Println("STORED HASH:", user.Password)
	// fmt.Println("PASSWORD RECEIVED:", password)

	// Compare provided password with stored hash
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Put claims into the JWT
	claims := jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}