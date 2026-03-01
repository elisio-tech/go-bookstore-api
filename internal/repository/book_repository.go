package repository

import (
	"errors"
	"go-bookstore-api/database"
	"go-bookstore-api/internal/models"
)

func FindAllBooks() ([]models.Book, error) {
	var books []models.Book

	err := database.DB.Find(&books)
	if err.Error != nil {
		return nil, err.Error
	}
	return books, nil
}

func FindBookById(id uint) (models.Book, error) {
	var book models.Book

	err := database.DB.First(&book, id)
	if err.Error != nil {
		return models.Book{}, err.Error
	}
	return book, nil
}

func CreateBook(book models.Book) (models.Book, error) {
	newBook := database.DB.Create(&book)

	if newBook.Error != nil {
		return models.Book{}, newBook.Error
	}
	return book, nil
}

func Delete(id uint) error {
	result := database.DB.Delete(&models.Book{}, id)

	if result.RowsAffected == 0 {
		return errors.New("Book not found!")
	}
	return nil
}
