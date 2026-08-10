package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controller.GetAllTasks)
	router.GET("/tasks/:id", controller.GetTaskDetails)

	router.POST("/tasks", controller.AddTask)

	router.PUT("/tasks/:id", controller.UpdateTask)
	
	router.DELETE("/tasks/:id", controller.DeleteTask)

	router.POST("/register", controller.Register)

	return router
}