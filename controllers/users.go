package controllers

import (
	"net/http"

	"github.com/fadhilradh/gin-gorm/initializers"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	type Response struct {
		Id int
		Email string
		Created_at string
	}

	var response []Response
	result := initializers.DB.Raw("SELECT id, email, created_at FROM users").Scan(&response)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error when getting users",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users" : response,
	})

}