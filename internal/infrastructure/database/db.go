package database

import (
	"fmt"

	"go-bookstore-api/internal/config"
	"go-bookstore-api/internal/domain/entity"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Get().DBPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir o banco: %w", err)
	}

	err = DB.AutoMigrate(&entity.Book{}, &entity.User{})
	if err != nil {
		return nil, fmt.Errorf("falha na migração: %w", err)
	}
	return DB, nil
}