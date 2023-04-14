package main

import (
	"github.com/fadhilradh/gin-gorm/controllers"
	"github.com/fadhilradh/gin-gorm/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initialize() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	initialize()

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my paradise",
		})
	})
	server.Use(cors.Default())
	server.Run()

	server.POST("/register", controllers.Register)
	server.POST("/login", controllers.Login)
	server.GET("/users", controllers.GetUsers)
	server.Run()
}