package models

type Book struct {
	ID     uint    `gorm:"AUTO_INCREMENT" json:"id"`
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	Price  float32 `json:"price" binding:"required"`
}
