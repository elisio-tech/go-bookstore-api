package repository

import (
	"go-bookstore-api/internal/database"

	"gorm.io/gorm"
)

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormRepository() *GormBookRepository {
	return &GormBookRepository{db: database.DB}
}
