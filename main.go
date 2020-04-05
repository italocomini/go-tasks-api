package main

import (
	"github.com/gin-gonic/gin"
	"github.com/italocomini/tasks/controllers"
	"github.com/italocomini/tasks/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	return r
}


func main() {
	r := setupRouter()
	r.Run()
}