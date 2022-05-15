package dto

// RegisterDTO is used when registering a new user
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required,min=1" validate:"min=3"`
	Email    string `json:"email" binding:"required,email" form:"email" `
	Password string `json:"password" binding:"required,min=6" form:"password" `
}
