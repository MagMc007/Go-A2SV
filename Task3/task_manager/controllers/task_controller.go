package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"

	"task_manager/data"
	"task_manager/models"
)

type TaskController struct {
	taskService data.TaskServices
}

func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (t *TaskController) AddTask(c *gin.Context) {
	var newTask models.Task

	// is there an error
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	t.taskService.AddTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func (t *TaskController) GetAllTasks(c *gin.Context) {
	allTasks := t.taskService.GetAllTasks()

	c.IndentedJSON(http.StatusOK, allTasks)
}

func (t *TaskController) GetTaskDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

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

func (t *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

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


func (t *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

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