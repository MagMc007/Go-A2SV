package usecases

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"

	"task_manager/mocks"
	"task_manager/domain"
	"task_manager/infrastructure"
)


type UserUsecaseTestSuite struct {
	suite.Suite

	userRepository *mocks.MockUserRepository
	userUsecase    domain.UserUsecase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.userRepository = mocks.NewMockUserRepository(suite.T())
	suite.userUsecase = NewUserUsecase(suite.userRepository)
}

func (suite *UserUsecaseTestSuite) TestRegisterSuccess() {
	userID := primitive.NewObjectID()

	user := domain.User{
		ID:       userID,
		Username: "magtest",
		Password: "hello",
		Role:     "admin",
	}

	// Username does not exist, proceed
	suite.userRepository.
		EXPECT().
		GetByUsername(user.Username).
		Return(domain.User{}, mongo.ErrNoDocuments)

	// Register the user with a hashed password
	suite.userRepository.
		EXPECT().
		Register(mock.MatchedBy(func(u domain.User) bool {
			return u.ID == user.ID &&
				u.Username == user.Username &&
				u.Role == user.Role &&
				u.Password != user.Password
		})).
		Return(domain.User{
			ID:       user.ID,
			Username: user.Username,
			Password: "some-hashed-password",
			Role:     user.Role,
		}, nil)

	result, err := suite.userUsecase.Register(user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, result.ID)
	assert.Equal(suite.T(), user.Username, result.Username)
	assert.Equal(suite.T(), user.Role, result.Role)
	assert.NotEqual(suite.T(), "hello", result.Password)
}

func (suite *UserUsecaseTestSuite) TestRegisterError() {
	userID := primitive.NewObjectID()

	user := domain.User{
		ID:       userID,
		Username: "magtest",
		Password: "hello",
		Role:     "admin",
	}

	regError := errors.New("username already exists")

	// username already exists
	suite.userRepository.
		EXPECT().
		GetByUsername(user.Username).
		Return(user, nil)

	result, err := suite.userUsecase.Register(user)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), regError, err)
	assert.Equal(suite.T(), domain.User{}, result)
}


func (suite *UserUsecaseTestSuite) TestLoginSuccess() {
	userID := primitive.NewObjectID()

	loginRequest := domain.LoginRequest{
		Username: "magtest",
		Password: "qwerty123",
	}

	hashed, _ := infrastructure.HashPwd([]byte(loginRequest.Password))

	storedUser := domain.User{
		ID:       userID,
		Username: "magtest",
		Password: hashed,
		Role:     "admin",
	}

	

	suite.userRepository.
		EXPECT().
		GetByUsername(loginRequest.Username).
		Return(storedUser, nil)

	result, err := suite.userUsecase.Login(
		loginRequest.Username,
		loginRequest.Password,
	)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), result)
}

func (suite *UserUsecaseTestSuite) TestLoginPasswordMismatch() {
	userID := primitive.NewObjectID()

	loginRequest := domain.LoginRequest{
		Username: "magtest",
		Password: "wrongpassword",
	}

	storedUser := domain.User{
		ID:       userID,
		Username: "magtest",
		Password: "$2a$10$u6B41GfnqVot7fVs2NXt1.SxqNDmhTjnXmE5Wzz54a/IefwRLOF5G",
		Role:     "admin",
	}

	loginErr := errors.New("invalid username or password")

	suite.userRepository.
		EXPECT().
		GetByUsername(loginRequest.Username).
		Return(storedUser, nil)

	result, err := suite.userUsecase.Login(
		loginRequest.Username,
		loginRequest.Password,
	)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), loginErr, err)
	assert.Empty(suite.T(), result)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}