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

	route.Run(":8080")
}
