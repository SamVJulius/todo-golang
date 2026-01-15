package main

import (
	"todo/controller"
	"todo/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectDB()
	database.SyncDatabase()
}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the ToDo App",
		})
	})

	router.POST("/create_task", controller.CreateTask)
	router.GET("/task/:id", controller.GetTaskByID)
	router.GET("/tasks", controller.GetAllTasks)
	router.GET("/status/:status", controller.GetTasksByStatus)
	router.GET("/category/:category", controller.GetTasksByCategory)
	router.PUT("/update/:id", controller.UpdateTask)
	router.DELETE("/delete/:id", controller.DeleteTask)
	router.PUT("/complete_task/:id", controller.CompleteTask)
	router.PUT("/mark_complete/:id", controller.MarkTaskCompleted)

	router.Run()
}