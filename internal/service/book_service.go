package service

import (
	"go-bookstore-api/internal/models"
	"go-bookstore-api/internal/repository"
)

func GetAllBooks() ([]models.Book, error) {
	return repository.FindAllBooks()
}
