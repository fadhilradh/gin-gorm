package initializers

import "github.com/fadhilradh/gin-gorm/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}