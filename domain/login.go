package domain

//Login - Entity login with all members of the technology company,
//who in turn are the direct operators of the system.
type Login struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

//ValidLogin - Valid entity Login.
func ValidLogin(login *Login) (*Login, *Err) {
	if login.Login == "" {
		return nil, NewErr().WithCode("login10")
	}
	if login.Password == "" {
		return nil, NewErr().WithCode("login20")
	}
	return login, nil
}

//LoginRepository - Define all operations entity Login.
type LoginRepository interface {

	//Save - Save new login.
	Save(login Login) (id int, err *Err)

	//Update - Update password login.
	Update(login Login) *Err

	//FindByLogin - Find login with login.
	FindByLogin(login string) (Login, *Err)
}
