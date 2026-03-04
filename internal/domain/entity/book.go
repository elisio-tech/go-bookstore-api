package entity

type Book struct {
	ID     uint    `json:"id" gorm:"AUTO_INCREMENT"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
