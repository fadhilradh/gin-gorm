package controllers

import (
	"net/http"

	"github.com/fadhilradh/gin-gorm/initializers"
	"github.com/fadhilradh/gin-gorm/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
 
func Register(c *gin.Context) {
	var body struct {
		Email 		string
		Password 	string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request body is not valid or unreadable",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error when hashing password",
		})

		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user) 

	if(result.Error != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error when creating user",
		})

		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success" : "user created",
	})
}