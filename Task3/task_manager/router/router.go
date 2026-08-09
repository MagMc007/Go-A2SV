package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

func SetupRouter(controller *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controller.GetAllTasks)
	router.GET("/tasks/:id", controller.GetTaskDetails)

	router.POST("/tasks", controller.AddTask)

	router.PUT("/tasks/:id", controller.UpdateTask)
	
	router.DELETE("/tasks/:id", controller.DeleteTask)

	return router
}