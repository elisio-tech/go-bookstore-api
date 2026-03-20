package database

import (
	"fmt"
	"go-bookstore-api/internal/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open("data/bookstore.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir o banco: %w", err)
	}

	err = DB.AutoMigrate(&entity.Book{})
	if err != nil {
		return nil, fmt.Errorf("Falha na migracao: %w", err)
	}
	return DB, nil
}
