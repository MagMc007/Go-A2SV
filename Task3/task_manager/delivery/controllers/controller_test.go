package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"task_manager/domain"
	"task_manager/mocks"
)

type ControllerTestSuite struct {
	suite.Suite

	taskUseCase *mocks.MockTaskUsecase
	userUseCase *mocks.MockUserUsecase
	controller  *Controller
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.taskUseCase = mocks.NewMockTaskUsecase(suite.T())
	suite.userUseCase = mocks.NewMockUserUsecase(suite.T())

	suite.controller = NewController(
		suite.taskUseCase,
		suite.userUseCase,
	)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func (suite *ControllerTestSuite) TestAddTaskSuccess() {
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Learn Go Testing",
		Description: "Learn controllers",
		Status:      "pending",
	}

	suite.taskUseCase.
		EXPECT().
		AddTask(mock.AnythingOfType("domain.Task")).
		Return(task, nil)

	router := setupRouter()
	router.POST("/tasks", suite.controller.AddTask)

	body := `{
		"title": "Learn Go Testing",
		"description": "Learn controllers",
		"status": "pending"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "Learn Go Testing")
}

func (suite *ControllerTestSuite) TestAddTaskInvalidBody() {
	router := setupRouter()
	router.POST("/tasks", suite.controller.AddTask)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(`invalid json`),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "invalid request body")
}

func (suite *ControllerTestSuite) TestAddTaskError() {
	task := domain.Task{
		Title:       "Learn Go Testing",
		Description: "Learn controllers",
		Status:      "pending",
	}

	repositoryErr := errors.New("database error")

	suite.taskUseCase.
		EXPECT().
		AddTask(mock.AnythingOfType("domain.Task")).
		Return(domain.Task{}, repositoryErr)

	router := setupRouter()
	router.POST("/tasks", suite.controller.AddTask)

	body := `{
		"title": "Learn Go Testing",
		"description": "Learn controllers",
		"status": "pending"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "database error")

	_ = task
}

func (suite *ControllerTestSuite) TestGetAllTasksSuccess() {
	taskID := primitive.NewObjectID()

	tasks := []domain.Task{
		{
			ID:          taskID,
			Title:       "Learn Go",
			Description: "Testing",
			Status:      "pending",
		},
	}

	suite.taskUseCase.
		EXPECT().
		GetAllTasks().
		Return(tasks, nil)

	router := setupRouter()
	router.GET("/tasks", suite.controller.GetAllTasks)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "Learn Go")
}

func (suite *ControllerTestSuite) TestGetAllTasksError() {
	suite.taskUseCase.
		EXPECT().
		GetAllTasks().
		Return(nil, errors.New("database error"))

	router := setupRouter()
	router.GET("/tasks", suite.controller.GetAllTasks)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "failed to get all tasks")
}

func (suite *ControllerTestSuite) TestGetTaskDetailsSuccess() {
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Learn Go",
		Description: "Testing",
		Status:      "pending",
	}

	suite.taskUseCase.
		EXPECT().
		GetTaskDetails(taskID).
		Return(task, nil)

	router := setupRouter()
	router.GET("/tasks/:id", suite.controller.GetTaskDetails)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/"+taskID.Hex(),
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "Learn Go")
}

func (suite *ControllerTestSuite) TestGetTaskDetailsInvalidID() {
	router := setupRouter()
	router.GET("/tasks/:id", suite.controller.GetTaskDetails)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/invalid-id",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "invalid task ID")
}

func (suite *ControllerTestSuite) TestGetTaskDetailsError() {
	taskID := primitive.NewObjectID()

	suite.taskUseCase.
		EXPECT().
		GetTaskDetails(taskID).
		Return(domain.Task{}, errors.New("task not found"))

	router := setupRouter()
	router.GET("/tasks/:id", suite.controller.GetTaskDetails)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/"+taskID.Hex(),
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "task not found")
}

func (suite *ControllerTestSuite) TestUpdateTaskSuccess() {
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated description",
		Status:      "completed",
	}

	suite.taskUseCase.
		EXPECT().
		UpdateTask(taskID, mock.AnythingOfType("domain.Task")).
		Return(task, nil)

	router := setupRouter()
	router.PUT("/tasks/:id", suite.controller.UpdateTask)

	body := `{
		"title": "Updated Task",
		"description": "Updated description",
		"status": "completed"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/"+taskID.Hex(),
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "Updated Task")
}

func (suite *ControllerTestSuite) TestUpdateTaskInvalidID() {
	router := setupRouter()
	router.PUT("/tasks/:id", suite.controller.UpdateTask)

	body := `{
		"title": "Updated Task"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/invalid-id",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "invalid task ID")
}

func (suite *ControllerTestSuite) TestUpdateTaskInvalidBody() {
	taskID := primitive.NewObjectID()

	router := setupRouter()
	router.PUT("/tasks/:id", suite.controller.UpdateTask)

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/"+taskID.Hex(),
		strings.NewReader(`invalid json`),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "improper task format")
}

func (suite *ControllerTestSuite) TestUpdateTaskError() {
	taskID := primitive.NewObjectID()

	suite.taskUseCase.
		EXPECT().
		UpdateTask(taskID, mock.AnythingOfType("domain.Task")).
		Return(domain.Task{}, errors.New("task not found"))

	router := setupRouter()
	router.PUT("/tasks/:id", suite.controller.UpdateTask)

	body := `{
		"title": "Updated Task",
		"status": "completed"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/"+taskID.Hex(),
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
}

func (suite *ControllerTestSuite) TestDeleteTaskSuccess() {
	taskID := primitive.NewObjectID()

	suite.taskUseCase.
		EXPECT().
		DeleteTask(taskID).
		Return(nil)

	router := setupRouter()
	router.DELETE("/tasks/:id", suite.controller.DeleteTask)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/tasks/"+taskID.Hex(),
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusNoContent, rec.Code)
}

func (suite *ControllerTestSuite) TestDeleteTaskInvalidID() {
	router := setupRouter()
	router.DELETE("/tasks/:id", suite.controller.DeleteTask)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/tasks/invalid-id",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "invalid task ID")
}

func (suite *ControllerTestSuite) TestDeleteTaskError() {
	taskID := primitive.NewObjectID()

	suite.taskUseCase.
		EXPECT().
		DeleteTask(taskID).
		Return(errors.New("task not found"))

	router := setupRouter()
	router.DELETE("/tasks/:id", suite.controller.DeleteTask)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/tasks/"+taskID.Hex(),
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "task not found")
}

func (suite *ControllerTestSuite) TestRegisterSuccess() {
	userID := primitive.NewObjectID()

	user := domain.User{
		ID:       userID,
		Username: "magtest",
		Password: "hashed-password",
		Role:     "admin",
	}

	suite.userUseCase.
		EXPECT().
		Register(mock.AnythingOfType("domain.User")).
		Return(user, nil)

	router := setupRouter()
	router.POST("/register", suite.controller.Register)

	body := `{
		"username": "magtest",
		"password": "hello",
		"role": "admin"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "magtest")
}

func (suite *ControllerTestSuite) TestLoginSuccess() {
	suite.userUseCase.
		EXPECT().
		Login("magtest", "qwerty123").
		Return("some-jwt-token", nil)

	router := setupRouter()
	router.POST("/login", suite.controller.Login)

	body := `{
		"username": "magtest",
		"password": "qwerty123"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "some-jwt-token")
}


func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}