package dto

// RegisterDTO is used when registering a new user
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=3"`
	Email    string `json:"email" binding:"required" form:"email" validate:"email"`
	Password string `json:"password" binding:"required" form:"password" validate:"min=6"`
}
