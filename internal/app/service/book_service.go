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
	book := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
		Price:  req.Price,
	}

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
	return s.bookRepo.List()
}

func (s *BookService) GetBookByID(id string) (*entity.Book, error) {
	return s.bookRepo.GetByID(id)
}

func (s *BookService) UpdateBook(id string, req dto.UpdateBookRequest) (*entity.Book, error) {
	existing, err := s.bookRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Author != "" {
		existing.Author = req.Author
	}
	if req.Price != 0 {
		existing.Price = req.Price
	}

	if err := s.bookRepo.Update(id, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *BookService) DeleteBook(id string) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("livro não encontrado")
		}
		return fmt.Errorf("falha ao excluir livro: %w", err)
	}
	return nil
}