package repository

import "go-bookstore-api/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByUsername(username string) (*entity.User, error)
}