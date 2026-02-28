package handlers

import (
	"go-bookstore-api/internal/models"
	"go-bookstore-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	books, err := service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error to get books",
		})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	var book models.Book

	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {}

func DeleteBook(ctx *gin.Context) {}

func UpdateBook(ctx *gin.Context) {}
