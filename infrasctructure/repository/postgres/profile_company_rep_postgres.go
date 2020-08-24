package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/domain/permission"
	"github.com/gabrielbo1/iroko/pkg"
)

type profileCompanyRepPostgres struct {
	tx *sql.Tx
}

//NewCompanyRepPostgres - Implementation domain.CompanyRepository with PostgreSQL.
func NewProfileCompanyRepPostgres(tx *sql.Tx) *profileCompanyRepPostgres {
	return &profileCompanyRepPostgres{tx: tx}
}

//Save - Save new ProfileCompany.
func (rep *profileCompanyRepPostgres) Save(profile permission.ProfileCompany) (int, *domain.Err) {
	return 0, nil
}

//Update - Update ProfileCompany.
func (rep *profileCompanyRepPostgres) Update(profile permission.ProfileCompany) *domain.Err {
	return nil
}

//FindById - Find ProfileCompany with id.
func (rep *profileCompanyRepPostgres) FindById(id int) (permission.ProfileCompany, *domain.Err) {
	return permission.ProfileCompany{}, nil
}

//FindByName - Profile page search with name.
func (rep *profileCompanyRepPostgres) FindByName(name string) (pkg.Page, *domain.Err) {
	return pkg.Page{}, nil
}
