package database

import (
	"errors"
	"go-bookstore-api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		return errors.New("Erro ao conetar ao banco de dados")
	}

	DB.AutoMigrate(&models.Book{})
	return nil
}
