package dto

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	User UserProfile `json:"user"`
}

type UserProfile struct{ 
	ID string `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
}