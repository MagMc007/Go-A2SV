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
	collection := database.Collection("tasks")

	service := data.NewTaskService(collection)
	controller := controllers.NewTaskController(service)

	r:= router.SetupRouter(controller)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}