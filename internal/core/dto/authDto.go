package core_dto

import "time"

type RegisterUserDTO struct {
	Username        string `json:"username" binding:"required,min=4,max=32"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6,max=48"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6,max=48,eqfield=Password"`
}

type RegisterUserResponse struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUserDto struct {
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=6,max=48"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ProfileUserResponse struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	RegisterAt time.Time `json:"register_at"`
}
