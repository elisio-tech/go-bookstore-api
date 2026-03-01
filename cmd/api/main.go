package main

import (
	"go-bookstore-api/database"
	"go-bookstore-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	route := gin.Default()
	route.GET("/books", handlers.GetBooks)
	route.GET("/books/:id", handlers.GetBook)
	route.POST("/books", handlers.CreateBook)
	route.DELETE("/books/:id", handlers.DeleteBook)

	route.Run(":8080")
}
