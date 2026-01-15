package database

import (
	"todo/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Task{})
}