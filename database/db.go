package database

import (
	"go-bookstore-api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Book{})
}
