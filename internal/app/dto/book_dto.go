package dto

type CreateBookRequest struct {
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	Price  float64 `json:"price" binding:"required,min=0"`
}

type BookResponse struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
