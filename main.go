package main

import (
	"fmt"

	"github.com/fadhilradh/gin-gorm/controllers"
	"github.com/fadhilradh/gin-gorm/initializers"
	"github.com/gin-gonic/gin"
)

func initialize() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	fmt.Println("Goodbye world !")
	initialize()

	server := gin.Default()
	server.POST("/register", controllers.Register)
	server.POST("/login", controllers.Login)
	server.Run()
}