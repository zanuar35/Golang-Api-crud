package dto

// UserUpdateDTO is used when client is updating user
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required" `
	Name     string `json:"name" form:"name" binding:"required" `
	Email    string `json:"email" form:"email" binding:"required" validate:"email" `
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min=6" `
}

// UserDTO is used when client is creating a new user
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required" `
	Email    string `json:"email" form:"email" binding:"required" validate:"email" `
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min=6" `
}
