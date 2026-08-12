package main

import (
	"log"
	"context"
	"time"


	"task_manager/delivery/controllers"
	"task_manager/usecases"
	"task_manager/repositories"
	"task_manager/data"
	"task_manager/delivery/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env")
	}

	client, err := data.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Printf("Error disconnecting MongoDB: %v", err)
        }
    }()

	database := client.Database("TaskManager")
	taskCollection := database.Collection("tasks")
	userCollection := database.Collection("users")

	// build repositories (talk to Mongo)
	taskRepo := repositories.NewTaskRepository(taskCollection)
	userRepo := repositories.NewUserRepository(userCollection)

	// build usecases (wrap repositories)
	taskService := usecases.NewTaskUsecase(taskRepo)
	userService := usecases.NewUserUsecase(userRepo)   
	
	controller := controllers.NewController(
		taskService,
		userService,
	)

	r:= router.SetupRouter(controller)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}