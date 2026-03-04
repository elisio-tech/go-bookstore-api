package main

import (
	"go-bookstore-api/internal/infrastructure/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	r.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"massage": "pong"})
		return
	})
	r.Run(":8080")
}
