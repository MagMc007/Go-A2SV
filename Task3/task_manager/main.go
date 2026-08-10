package main

import (
	"log"

	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

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