package domain

import "github.com/gabrielbo1/iroko/pkg"

//Company - Company Entity company with all related information,
//and the other relationships belong to a company and technology
//company in turn have many client companies.
type Company struct {
	ID      int    `json:"id"`
	Version int    `json:"version"`
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
}

//ValidCompany - Valid Company entity.
func ValidCompany(company *Company) (*Company, *Err) {
	if company.Name == "" {
		return nil, NewErr().WithCode("company10")
	}
	return company, nil
}

//CompanyRepository - Define all operations entity Company.
type CompanyRepository interface {

	//Save - Save new company entity.
	Save(company Company) (id int, err *Err)

	//Update - Update company entity.
	Update(company Company) *Err

	//FindById - Find company with ID
	FindById(id int) (Company, *Err)

	//FindByUUID - Find company with UUID.
	FindByUUID(uuid string) (Company, *Err)

	//FindByPage - Find by company with name, pageable query.
	FindByPage(name string, page pkg.Page) (pkg.Page, *Err)
}
