package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"task_manager/data"
	"task_manager/models"
)

type Controller struct {
	taskService data.TaskServices
	userService data.UserServices
}

func NewController(
	taskService *data.TaskService,
	userService *data.UserService,
	) *Controller {
	return &Controller{
		taskService: taskService,
		userService: userService,
	}
}

func (t *Controller) AddTask(c *gin.Context) {
	var newTask models.Task

	// is there an error
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	createdTask, err := t.taskService.AddTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdTask)
}

func (t *Controller) GetAllTasks(c *gin.Context) {
	allTasks, err := t.taskService.GetAllTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, allTasks)
}

func (t *Controller) GetTaskDetails(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	task, err := t.taskService.GetTaskDetails(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (t *Controller) UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	var task models.Task
	
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "improper task format",
		})
		return
	}

	updatedTask, err := t.taskService.UpdateTask(id, task)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}


func (t *Controller) DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	err = t.taskService.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return 
	}

	c.Status(http.StatusNoContent)
}


func (t * Controller) Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
	}

	// set default user role to user(normal one)
	newUser.Role = "user"

	createdUser, err := t.userService.Register(newUser)

	if err != nil {
		if err.Error() == "username already exists" {
			c.IndentedJSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdUser)
} 