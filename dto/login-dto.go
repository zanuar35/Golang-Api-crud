package dto

// LoginDTO is used when client is logging in
type LoginDTO struct {
	Email    string `json:"email" binding:"required,min=1" form:"email" `
	Password string `json:"password" binding:"required,min=6" form:"password" `
}
