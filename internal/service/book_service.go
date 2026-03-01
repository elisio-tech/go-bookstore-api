package service

import (
	"errors"
	"go-bookstore-api/internal/models"
	"go-bookstore-api/internal/repository"
)

func GetAllBooks() ([]models.Book, error) {
	return repository.FindAllBooks()
}

func GetBookById(id int) (models.Book, error) {
	return repository.FindBookById(uint(id))
}

func SaveBook(book models.Book) (models.Book, error) {
	if book.Author == "" {
		return models.Book{}, errors.New("O autor é obrigatório")
	}
	if book.Title == "" {
		return models.Book{}, errors.New("O titulo é obrigatório")
	}

	if book.Price <= 0 {
		return models.Book{}, errors.New("preço inválido")
	}
	return repository.CreateBook(book)
}

func DeleteBookByID(id uint) error {
	return repository.Delete(id)
}
