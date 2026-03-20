package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     string  `gorm:"primaryKey" json:"id"`
	Title  string  `json:"title"`
	Author string  `josn:"author"`
	Price  float64 `json:"price"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = "bk_" + uuid.NewString()
	}
	return nil
}
