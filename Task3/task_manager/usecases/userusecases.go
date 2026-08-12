package usecases

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"task_manager/domain"
	"task_manager/infrastructure"
)

type UserUseCase struct {
    userRepository domain.UserRepository
}

func NewUserUseCase(collection *mongo.Collection) *UserUseCase {
	return &UserUserCase{
		userRepository: userRepository,
	}
}

func (ur *UserUsecase) Register(user domain.User) (domain.User, error) {
	// lookup username
	_, err := uu.userRepository.GetByUsername(user.Username)
	
	if err == nil {
		return domain.User{}, errors.New("username already exists")
	}

	if err != mongo.ErrNoDocuments {
		return domain.User{}, err
	}

	// hash pwd
	hashedPassword, err := infrastructure.HashPwd([]byte(user.Password))
	if err != nil {
		return domain.User{}, err
	}

	user.Password = hashedPassword

	return uu.userRepository.Register(user)
}

func (uu *UserUsecase) Login(username string, password string) (string, error) {
	// look up username
	user, err := uu.userRepository.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// check password against stored hash
	if !infrastructure.CheckPwd(user.Password, password) {
		return "", errors.New("invalid username or password")
	}

	// generate token with claims
	token, err := infrastructure.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}