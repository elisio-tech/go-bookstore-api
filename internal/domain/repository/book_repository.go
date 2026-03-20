package repository

import (
	"go-bookstore-api/internal/domain/entity"
)

// Interface define o contrato. A implementação fica na Infrastructure.
type BookRepository interface {
	Create(book *entity.Book) error
	GetByID(id string) (*entity.Book, error)
	List() ([]entity.Book, error)
	Update(id string, book *entity.Book) error
	Delete(id string) error
}
