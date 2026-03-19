package database

import (
	"fmt"
	"go-bookstore-api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir o banco: %w", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, fmt.Errorf("Falha na migracao: %w", err)
	}
	return db, nil
}
