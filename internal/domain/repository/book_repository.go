package repository

import (
	"go-bookstore-api/internal/app/dto"
	"go-bookstore-api/internal/domain/entity"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	GetByID(id string) (*entity.Book, error)
	Create(book *entity.Book) (dto.RequestBook, error)
	Update(book *entity.Book) error
	Delete(id string) error
}
