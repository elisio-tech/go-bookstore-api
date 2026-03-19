package repository

import (
	"gorm.io/gorm"
)

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}
