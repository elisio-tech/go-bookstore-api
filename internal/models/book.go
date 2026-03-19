package models

type Book struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `josn:"author"`
	Price  int64  `json:"price"`
}
