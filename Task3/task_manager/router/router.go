package router

import (
	"task_manager/controllers"

	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()

	protected := router.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())

	protected.GET(
		"",
		middleware.AdminOnly(),
		controller.GetAllTasks,
	)

	protected.GET("/:id", controller.GetTaskDetails)
	protected.POST("", controller.AddTask)
	protected.PUT("/:id", controller.UpdateTask)
	protected.DELETE("/:id", controller.DeleteTask)

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	return router
}