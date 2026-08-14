package usecases

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"task_manager/mocks"
	"task_manager/domain"
)

type TaskUsecaseTestSuite struct {
	suite.Suite

	taskRepository *mocks.MockTaskRepository
	taskUsecase    domain.TaskUsecase
}

func (suite *TaskUsecaseTestSuite) SetupTest() {
	suite.taskRepository = mocks.NewMockTaskRepository(suite.T())
	suite.taskUsecase = NewTaskUsecase(suite.taskRepository)
}

func (suite *TaskUsecaseTestSuite) TestAddTaskSuccess() {
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Learn Go Testing",
		Description: "Learn Testify and Mockery",
		Status:      "pending",
	}

	suite.taskRepository.
		EXPECT().
		AddTask(task).
		Return(task, nil)

	result, err := suite.taskUsecase.AddTask(task)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, result)
	assert.Equal(suite.T(), taskID, result.ID)
}

func (suite *TaskUsecaseTestSuite) TestGetAllTasksSuccess() {
	// add a task
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Learn Go Testing",
		Description: "Learn Testify and Mockery",
		Status:      "pending",
	}

	tasks := []domain.Task{
		task,
	}
	
	// get all tasks
	// we expect an array
	suite.taskRepository.
		EXPECT().
		GetAllTasks().
		Return(tasks, nil)

	result, err := suite.taskUsecase.GetAllTasks()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), tasks, result)
}

func (suite *TaskUsecaseTestSuite) TestGetTaskDetailsSuccess() {
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Learn Go Testing",
		Description: "Learn Testify and Mockery",
		Status:      "pending",
	}

	suite.taskRepository.
		EXPECT().
		GetTaskDetails(taskID).
		Return(task, nil)

	result, err := suite.taskUsecase.GetTaskDetails(taskID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, result)
	assert.Equal(suite.T(), taskID, result.ID)
}

func (suite *TaskUsecaseTestSuite) TestGetTaskDetailsError() {
	taskID := primitive.NewObjectID()

	repositoryErr := errors.New("invalid task ID")

	suite.taskRepository.
		EXPECT().
		GetTaskDetails(taskID).
		Return(domain.Task{}, repositoryErr)

	result, err := suite.taskUsecase.GetTaskDetails(taskID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), repositoryErr, err)
	assert.Equal(suite.T(), domain.Task{}, result.ID)
}

func (suite *TaskUsecaseTestSuite) TestUpdateTaskSuccess() {
	taskID := primitive.NewObjectID()

	task := domain.Task{
		ID:          taskID,
		Title:       "Learn Go Testing",
		Description: "Learn Testify and Mockery",
		Status:      "Completed",
	}

	updatedTask := domain.Task{
		Title:       "Learn Go Testing",
		Description: "Learn Testify and Mockery",
		Status:      "Completed",
	}	


	suite.taskRepository.
		EXPECT().
		UpdateTask(taskID, updatedTask).
		Return(task, nil)
	
	result, err := suite.taskUsecase.UpdateTask(taskID, updatedTask)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, result)
	assert.Equal(suite.T(), taskID, result.ID)
}

func (suite *TaskUsecaseTestSuite) TestUpdateTaskError() {
	taskID := primitive.NewObjectID()

	updatedTask := domain.Task{
		Title:       "Learn Go Testing",
		Description: "Learn Testify and Mockery",
		Status:      "Completed",
	}

	repositoryErr := errors.New("invalid task ID")

	suite.taskRepository.
		EXPECT().
		UpdateTask(taskID, updatedTask).
		Return(domain.Task{}, repositoryErr)

	result, err := suite.taskUsecase.UpdateTask(taskID, updatedTask)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), repositoryErr, err)
	assert.Equal(suite.T(), domain.Task{}, result)
}

func (suite *TaskUsecaseTestSuite) TestDeleteTaskSuccess() {
	taskID := primitive.NewObjectID()

	suite.taskRepository.
		EXPECT().
		DeleteTask(taskID).
		Return(nil)
	
	err := suite.taskUsecase.DeleteTask(taskID)

	assert.NoError(suite.T(), err)
}	

func (suite *TaskUsecaseTestSuite) TestDeleteTaskError() {
	taskID := primitive.NewObjectID()

	repositoryErr := errors.New("invalid task ID")

	suite.taskRepository.
		EXPECT().
		DeleteTask(taskID).
		Return(repositoryErr)
	
		err := suite.taskUsecase.DeletedTask(taskID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), err, repositoryErr)
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}