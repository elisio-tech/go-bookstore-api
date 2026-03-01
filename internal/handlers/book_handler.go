package handlers

import (
	"go-bookstore-api/internal/models"
	"go-bookstore-api/internal/service"
	"go-bookstore-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	books, err := service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro a obter os livros",
		})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	var book models.Book
	id, err := utils.StringToInt(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "O ID deve ser um número inteiro"})
		return
	}

	book, err = service.GetBookById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "O corpo da requisição é inválido ou mal formatado",
		})
		return
	}

	newBook, err := service.SaveBook(newBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, newBook)
}

func DeleteBook(ctx *gin.Context) {
	id, err := utils.StringToInt(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "O ID deve ser um número inteiro"})
		return
	}

	err = service.DeleteBookByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
