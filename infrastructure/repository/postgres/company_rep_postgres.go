package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/pkg"
)

type companyRepPostgres struct {
	tx *sql.Tx
}

//NewCompanyRepPostgres - Implementation domain.CompanyRepository with PostgreSQL.
func NewCompanyRepPostgres(tx *sql.Tx) *companyRepPostgres {
	return &companyRepPostgres{tx: tx}
}

//Save - Save new company entity.
func (rep *companyRepPostgres) Save(company domain.Company) (id int, err *domain.Err) {
	return 0, nil
}

//Update - Update company entity.
func (rep *companyRepPostgres) Update(company domain.Company) *domain.Err {
	return nil
}

//FindById - Find company with ID.
func (rep *companyRepPostgres) FindById(id int) (domain.Company, *domain.Err) {
	return domain.Company{}, nil
}

//FindByUUID - Find company with UUID.
func (rep *companyRepPostgres) FindByUUID(uuid string) (domain.Company, *domain.Err) {
	return domain.Company{}, nil
}

//FindByPage - Find by company with name, pageable query.
func (rep *companyRepPostgres) FindByPage(name string, page pkg.Page) (pkg.Page, *domain.Err) {
	return pkg.Page{}, nil
}
