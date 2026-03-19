package repository

import "go-bookstore-api/internal/models"

type BookRepository interface {
	FindAll() ([]models.Book, error)
	GetByID(id string) (*models.Book, error)
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(id string) error
}
