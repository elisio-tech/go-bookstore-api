package repository

import (
	"go-bookstore-api/internal/domain/entity"
	"go-bookstore-api/internal/infrastructure/database"

	"gorm.io/gorm"
)

type SQLiteBookRepository struct {
	db *gorm.DB
}

// lida com banco de dados abre e fecha conexao
func NewSQLiteBookRepository() *SQLiteBookRepository {
	return &SQLiteBookRepository{db: database.DB}
}

func (r *SQLiteBookRepository) List() ([]entity.Book, error) {
	var books []entity.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *SQLiteBookRepository) Create(book *entity.Book) error {
	return r.db.Create(book).Error
}

func (r *SQLiteBookRepository) GetByID(id string) (*entity.Book, error) {
	var book entity.Book

	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &book, err
}

func (r *SQLiteBookRepository) Update(id string, book *entity.Book) error {
	return r.db.Model(&entity.Book{}).Where("id = ?", id).Updates(book).Error
}

func (r *SQLiteBookRepository) Delete(id string) error {
	return r.db.Delete(&entity.Book{}, id).Error
}
