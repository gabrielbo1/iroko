package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/pkg"
)

type loginRepPostgres struct {
	tx *sql.Tx
}

//NewLoginRepPostgres - Implementation domain.LoginRepository with PostgreSQL.
func NewLoginRepPostgres(tx *sql.Tx) *loginRepPostgres {
	return &loginRepPostgres{tx: tx}
}

//Save - Save new login.
func (rep *loginRepPostgres) Save(login domain.Login) (id int, err *domain.Err) {
	return 0, nil
}

//Update - Update password login.
func (rep *loginRepPostgres) Update(login domain.Login) *domain.Err {
	return nil
}

//FindByLogin - Find login with login.
func (rep *loginRepPostgres) FindByLogin(login string) (domain.Login, *domain.Err) {
	return domain.Login{}, nil
}

//FindByPage - Find by Login with name, pageable query.
func (rep *loginRepPostgres) FindByPage(name string, page pkg.Page) (pkg.Page, *domain.Err) {
	return pkg.Page{}, nil
}
