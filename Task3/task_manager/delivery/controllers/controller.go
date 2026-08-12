package	controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"task_manager/domain"
)

type Controller struct {
	taskUseCase domain.TaskUsecase
	userUseCase domain.UserUsecase
}

func NewController(taskUseCase domain.TaskUsecase, userUseCase domain.UserUsecase) *Controller {
	return &Controller{
		taskUseCase: taskUseCase,
		userUseCase: userUseCase,
	}
}

func (t *Controller) AddTask(c *gin.Context) {
	var newTask domain.Task

	// is there an error
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	createdTask, err := t.taskUseCase.AddTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdTask)
}

func (t *Controller) GetAllTasks(c *gin.Context) {
	allTasks, err := t.taskUseCase.GetAllTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all tasks",
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

	task, err := t.taskUseCase.GetTaskDetails(id)
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

	var task domain.Task
	
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "improper task format",
		})
		return
	}

	updatedTask, err := t.taskUseCase.UpdateTask(id, task)

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

	err = t.taskUseCase.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return 
	}

	c.Status(http.StatusNoContent)
}


func (t * Controller) Register(c *gin.Context) {
	var newUser domain.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// set default user role to user(normal one)
	newUser.Role = "user"

	createdUser, err := t.userUseCase.Register(newUser)

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

func (t *Controller) Login(c *gin.Context) {
    var loginRequest domain.LoginRequest

    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "invalid request body",
        })
        return
    }

    token, err := t.userUseCase.Login(
        loginRequest.Username,
        loginRequest.Password,
    )

    if err != nil {
        if err.Error() == "invalid username or password" {
            c.IndentedJSON(http.StatusUnauthorized, gin.H{
                "error": err.Error(),
            })
            return
        }

        c.IndentedJSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{
        "token": token,
    })
}