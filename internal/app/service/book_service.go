package service

import (
	"errors"
	"fmt"
	"go-bookstore-api/internal/app/dto"
	"go-bookstore-api/internal/domain/entity"
	"go-bookstore-api/internal/domain/repository"

	"gorm.io/gorm"
)

type BookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{bookRepo: repo}
}

func (s *BookService) CreateBook(req dto.RequestBook) (*dto.ResponseBook, error) {
	if req.Price < 0 {
		return nil, fmt.Errorf("preço deve ser maior que zero")
	}

	book := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
		Price:  req.Price,
	}

	// chama o repository
	if err := s.bookRepo.Create(book); err != nil {
		return nil, err
	}

	return &dto.ResponseBook{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}, nil
}

func (s *BookService) GetBooks() ([]entity.Book, error) {
	books, err := s.bookRepo.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) GetBookByID(id string) (*entity.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) UpdateBook(id string, book *entity.Book) error {
	return s.bookRepo.Update(id, book)
}

func (s *BookService) DeleteBook(id string) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Livro nao encontrado")
		}
		return fmt.Errorf("failed to delete book: %w", err)
	}
	return nil
}
