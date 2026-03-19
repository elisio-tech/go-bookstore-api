package repository

import (
	"go-bookstore-api/internal/models"

	"gorm.io/gorm"
)

type GormBookRepository struct {
	db *gorm.DB
}

// lida com banco de dados abre e fecha conexao
func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}

func (r *GormBookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *GormBookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *GormBookRepository) GetByID(id string) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	return &book, err
}

func (r *GormBookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *GormBookRepository) Delete(id string) error {
	return r.db.Delete(&models.Book{}, id).Error
}
