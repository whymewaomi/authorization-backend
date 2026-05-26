package domain

import "time"

type User struct {
	ID int
	Username string 
	Email string
	Password string
	PasswordConfirm string
	RegisterAt time.Time
}

type UserToken struct {
	ID int
	AccessToken string
}

func NewLoginUser(
	Username string,
	Password string,
) *User {
	return &User{
		Username: Username,
		Password: Password,
	}
}

func NewRegisterUser(
	Username string,
	Email string,
	Password string,
	PasswordConfirm string,
) *User {
	return &User{
		Username: Username,
		Email: Email,
		Password: Password,
		PasswordConfirm: PasswordConfirm,
	}
}

func NewUserToken(
	ID int,
	AccessToken string,
) *UserToken {
	return &UserToken{
		ID: ID,
		AccessToken: AccessToken,
	}
}