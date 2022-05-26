package dto

// BookUpdateDTO is a model that is used when client is updating a book
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required" `
	Title       string `json:"title" form:"title" binding:"required" `
	Description string `json:"description" form:"description" binding:"required" `
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty" `
}

// BookCreatedDTO is a model that is used when client is creating a new book
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required" `
	Description string `json:"description" form:"description" binding:"required" `
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty" `
}
