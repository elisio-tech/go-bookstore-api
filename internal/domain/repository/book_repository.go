package repository

import "go-bookstore-api/internal/domain/entity"

// Interface define o contrato. A implementação fica na Infrastructure
type BookRepository interface {
	Create(book *entity.Book) (*entity.Book, error)
	GetByID(id uint) (*entity.Book, error)
	List() ([]entity.Book, error)
	Update(id uint, book *entity.Book) (*entity.Book, error)
	Delete(id uint) error
}
