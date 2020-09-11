package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/pkg"
)

type loginCompanyRepPostgres struct {
	tx *sql.Tx
}

//NewLoginCompanyRepPostgres - Implementation domain.LoginCompanyRepository with PostgreSQL.
func NewLoginCompanyRepPostgres(tx *sql.Tx) *loginCompanyRepPostgres {
	return &loginCompanyRepPostgres{tx: tx}
}

//Save - Save LoginCompany entity.
func (rep *loginCompanyRepPostgres) Save(login domain.LoginCompany) (id int, err *domain.Err) {
	return 0, nil
}

//Update - Update LoginCompany entity.
func (rep *loginCompanyRepPostgres) Update(login domain.LoginCompany) *domain.Err {
	return nil
}

//FindByLogin - Find LoginCompany with login.
func (rep *loginCompanyRepPostgres) FindByLogin(login string) (id int, err *domain.Err) {
	return 0, nil
}

//FindByPage - Find by Login with name, pageable query.
func (rep *loginCompanyRepPostgres) FindByPage(name string, page pkg.Page) (pkg.Page, *domain.Err) {
	return pkg.Page{}, nil
}
