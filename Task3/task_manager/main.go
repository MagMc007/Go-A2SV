package main

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	service := data.NewTaskService()
	controller := controllers.NewTaskController(service)

	r:= router.SetupRouter(controller)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}