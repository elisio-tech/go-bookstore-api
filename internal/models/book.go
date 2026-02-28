package models

type Book struct {
	ID     uint    `gorm:"AUTO_INCREMENT" json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}
