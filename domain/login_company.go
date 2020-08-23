package domain

import "github.com/gabrielbo1/iroko/pkg"

//LoginCompany - Entity for access control of persons associated
//with client companies that may eventually access external
//functionalities by system or integrations by api.
type LoginCompany struct {
	ID       int     `json:"id"`
	Version  int     `json:"version"`
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Company  Company `json:"company"`
	Email    string  `json:"email"`
}

//ValidLoginCompany - Valid LoginCompany entity.
func ValidLoginCompany(login *LoginCompany) (*LoginCompany, *Err) {
	if login.Login == "" {
		return nil, NewErr().WithCode("logincompany10")
	}
	if login.Password == "" {
		return nil, NewErr().WithCode("logincompany20")
	}
	if login.Company.ID == 0 {
		return nil, NewErr().WithCode("logincompany30")
	}
	if login.Email == "" || !isEmailValid(login.Email) {
		return nil, NewErr().WithCode("logincompany40")
	}
	return login, nil
}

//LoginCompanyRepository - All operations
//in entity LoginCompanyRepository.
type LoginCompanyRepository interface {

	//Save - Save LoginCompany entity.
	Save(login LoginCompany) (id int, err *Err)

	//Update - Update LoginCompany entity.
	Update(login LoginCompany) *Err

	//FindByLogin - Find LoginCompany with login.
	FindByLogin(login string) (id int, err *Err)

	//FindByPage - Find by company with name, pageable query.
	FindByPage(name string, page pkg.Page) (pkg.Page, *Err)
}
