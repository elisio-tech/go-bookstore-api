package dto

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token    string `json:"token"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type UpdateBookRequest struct {
	Title  string  `json:"title" binding:"omitempty,min=1,max=255"`
	Author string  `json:"author" binding:"omitempty,min=1,max=255"`
	Price  float64 `json:"price" binding:"omitempty,min=0"`
}