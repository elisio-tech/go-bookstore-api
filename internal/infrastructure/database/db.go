package database

import (
	"go-bookstore-api/internal/domain/entity"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conetar ao banco de dados:", err)
	}

	err = DB.AutoMigrate(&entity.Book{})
	if err != nil {
		log.Fatal("Falha ao migrar tabelas:", err)
	}
}
