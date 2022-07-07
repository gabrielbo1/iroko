package domain

import "github.com/gabrielbo1/iroko/pkg"

// Subsidiary - Subsidiary within the permissions control, basically
// the company and the systems associated with it are registered,
// from there the logins associated with that company with the
// respective permissions.
// Being this abstraction to compose multi-company cases,
// being a company for many subsidiaries.
type Subsidiary struct {
	CompanyId string   `json:"company_id"`
	Name      string   `json:"name"`
	Systems   []System `json:"systems"`
	Users     []User   `json:"users"`
}

// NewSubsidiary - Creare new subsidiary valid.
func NewSubsidiary(sub Subsidiary) (*Subsidiary, *pkg.Err) {
	if !pkg.UuidIsValid(sub.CompanyId) {
		return nil, pkg.NewErr().WithCode("PERMISSION_SUBSIDIARY_10")
	}
	if !pkg.NameIsValid(sub.Name) {
		return nil, pkg.NewErr().WithCode("PERMISSION_SUBSIDIARY_20")
	}
	return &Subsidiary{
		CompanyId: sub.CompanyId,
		Name:      sub.Name,
		Systems:   sub.Systems,
		Users:     sub.Users,
	}, nil
}

// Company - Company entity that defines the global ID of an organization within the Iroko systems,
// being an Id for the entire company and a second Id for each subsidiary of
// the company, the login structure and permissions control were linked to a subsidiary,
// that is, a system and a user belong to a subsidiary which in turn belongs to a company.
// Logins linked directly to the company are considered administrator-type logins,
// they have access to general system administration routines.
type Company struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Subsidiaries []Subsidiary `json:"subsidiaries"`
	Users        []User       `json:"users"`
}

//NewCompany - Create new valid company.
func NewCompany(company Company) (*Company, *pkg.Err) {
	if !pkg.NameIsValid(company.Name) {
		return nil, pkg.NewErr().WithCode("PERMISSION_COMPANY_10")
	}
	return &Company{
		Name:         company.Name,
		Subsidiaries: company.Subsidiaries,
		Users:        company.Users,
	}, nil
}
