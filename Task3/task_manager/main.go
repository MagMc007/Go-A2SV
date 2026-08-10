package main

import (
	"log"

	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	client, err := data.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("TaskManager")
	taskCollection := database.Collection("tasks")
	userCollection := database.Collection("users")

	taskService := data.NewTaskService(taskCollection)
	userService := data.NewUserService(userCollection)

	controller := controllers.NewController(
		taskService,
		userService,
	)

	r:= router.SetupRouter(controller)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}