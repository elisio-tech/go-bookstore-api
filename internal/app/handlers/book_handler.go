package handlers

import (
	"errors"
	"net/http"
	"strings"

	"go-bookstore-api/internal/app/dto"
	"go-bookstore-api/internal/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.CreateBook(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.service.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	book, err := h.service.GetBookByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "livro não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.UpdateBook(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "livro não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id é obrigatório"})
		return
	}

	err := h.service.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "livro excluído com sucesso"})
}