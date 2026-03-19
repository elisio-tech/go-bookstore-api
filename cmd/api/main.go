package main

import (
	"go-bookstore-api/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	r.Run()
}
