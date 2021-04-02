package domain

import (
	"regexp"
)

//Login - Entity login with all members of the technology company,
//who in turn are the direct operators of the system.
type Login struct {
	ID       int    `json:"id"`
	Version  int    `json:"version"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//ValidLogin - Valid entity Login.
func (login *Login) ValidLogin() (*Login, *Err) {
	if login.Login == "" {
		return nil, NewErr().WithCode("login10")
	}
	if login.Password == "" {
		return nil, NewErr().WithCode("login20")
	}
	if login.Email == "" || !isEmailValid(login.Email) {
		return nil, NewErr().WithCode("login30")
	}
	return login, nil
}

//Validate email regex based on W3C pattern.
//https://golangcode.com/validate-an-email-address/
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
//https://golangcode.com/validate-an-email-address/
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

//LoginRepository - Define all operations entity Login.
type LoginRepository interface {

	//Save - Save new Login.
	Save(login Login) (id int, err *Err)

	//Update - Update password Login.
	Update(login Login) *Err

	//FindByLogin - Find Login with login.
	FindByLogin(id int) (Login, *Err)

	//Delete - Delete login record with login.
	Delete(id int) *Err
}
