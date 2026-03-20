package main

import (
	"go-bookstore-api/internal/app/handlers"
	"go-bookstore-api/internal/app/service"

	"go-bookstore-api/internal/infrastructure/database"
	"go-bookstore-api/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	// repositorio
	repo := repository.NewSQLiteBookRepository()

	// service
	bookService := service.NewBookService(repo)

	// handler
	bookHandler := handlers.NewBookHandler(bookService)

	r := gin.Default()

	r.GET("/books", bookHandler.GetBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
	r.POST("books/:id", bookHandler.UpdateBook)
	r.POST("/books", bookHandler.CreateBook)

	r.Run()
}
