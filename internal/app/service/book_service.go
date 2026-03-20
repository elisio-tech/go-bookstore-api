package service

import (
	"fmt"
	"go-bookstore-api/internal/app/dto"
	"go-bookstore-api/internal/domain/entity"
	"go-bookstore-api/internal/domain/repository"
)

type BookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{bookRepo: repo}
}

func (s *BookService) CreateBook(request dto.RequestBook) (*dto.ResponseBook, error) {
	if request.Price < 0 {
		return nil, fmt.Errorf("preço não pode ser negativo")
	}

	book := &entity.Book{
		Title:  request.Title,
		Author: request.Author,
		Price:  request.Price,
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
		return nil, fmt.Errorf("Livro nao encontrado", err)
	}
	return book, nil
}

func (s *BookService) UpdateBook(id string, book *entity.Book) error {
	return s.bookRepo.Update(id, book)
}

func (s *BookService) DeleteBook(id string) error {
	return s.bookRepo.Delete(id)
}
