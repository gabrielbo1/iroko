package domain

import (
	"net/mail"

	"github.com/gabrielbo1/iroko/pkg"
)

// User basic informations.
type User struct {
	Id       string `json:"id"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser - Create new user and validate informations.
func NewUser(user User) (*User, *pkg.Err) {
	if !pkg.NameIsValid(user.Nick) {
		return nil, pkg.NewErr().WithCode("PERMISSION_USER_10")
	}
	if !pkg.NameIsValid(user.Email) {
		return nil, pkg.NewErr().WithCode("PERMISSION_USER_20")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return nil, pkg.NewErr().WithCode("PERMISSION_USER_20")
	}
	if !pkg.NameIsValid(user.Password) {
		return nil, pkg.NewErr().WithCode("PERMISSION_USER_30")
	}
	return &User{
		Nick:     user.Nick,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
