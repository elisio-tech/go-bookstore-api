package repository

import (
	"go-bookstore-api/internal/domain/entity"
	"go-bookstore-api/internal/infrastructure/database"

	"gorm.io/gorm"
)

type SQLiteUserRepository struct {
	db *gorm.DB
}

func NewSQLiteUserRepository() *SQLiteUserRepository {
	return &SQLiteUserRepository{db: database.DB}
}

func (r *SQLiteUserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *SQLiteUserRepository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}