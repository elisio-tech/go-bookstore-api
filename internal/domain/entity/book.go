package entity

type Book struct {
	ID     string  `gorm:"primaryKey" json:"id"`
	Title  string  `json:"title"`
	Author string  `josn:"author"`
	Price  float64 `json:"price"`
}
