package router

import(
	"task_manager/infrastructure"
	"task_manager/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()

	protected := router.Group("/tasks")
	protected.Use(infrastructure.AuthMiddleware())

	protected.GET(
		"",
		infrastructure.AdminOnly(),
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