package handlers

import (
	"go-bookstore-api/internal/app/dto"
	"go-bookstore-api/internal/app/service"
	"go-bookstore-api/internal/domain/entity"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(svc *service.BookService) *BookHandler {
	return &BookHandler{service: svc}
}

func (h *BookHandler) CreateBook(ctx *gin.Context) {
	var req dto.RequestBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.CreateBook(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.service.GetBooks()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, books)
}

func (h *BookHandler) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	book, err := h.service.GetBookByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, book)
}

func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")

	// 2. Faz o Bind do JSON para o DTO de Request
	var req entity.Book
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	err := h.service.UpdateBook(id, &req)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err})
	}

	ctx.JSON(200, req)
}

func (h *BookHandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Book deleted successfully"})
}
