package repository

import (
	"go-bookstore-api/database"
	"go-bookstore-api/internal/models"
)

func FindAllBooks() ([]models.Book, error) {
	var books []models.Book

	err := database.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
