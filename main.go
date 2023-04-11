package main

import (
	"fmt"

	"github.com/fadhilradh/gin-gorm/initializers"
	"github.com/gin-gonic/gin"
)

func initialize() {
	initializers.LoadEnv()
}

func main() {
	fmt.Println("Goodbye world !")
	initialize()

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run()
}