package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/domain/permission"
	"github.com/gabrielbo1/iroko/pkg"
)

type profileRepPostgres struct {
	tx *sql.Tx
}

//NewProfileRepPostgres - Implementation permission.ProfileRepository with PostgreSQL.
func NewProfileRepPostgres(tx *sql.Tx) *profileRepPostgres {
	return &profileRepPostgres{tx: tx}
}

//Save - Save new Profile.
func (rep *profileRepPostgres) Save(profile permission.Profile) (int, *domain.Err) {
	return 0, nil
}

//Update - Update Profile.
func (rep *profileRepPostgres) Update(profile permission.Profile) *domain.Err {
	return nil
}

//FindById - Find Profile with id.
func (rep *profileRepPostgres) FindById(id int) (permission.Profile, *domain.Err) {
	return permission.Profile{}, nil
}

//FindByName - Profile page search with name.
func (rep *profileRepPostgres) FindByName(name string) (pkg.Page, *domain.Err) {
	return pkg.Page{}, nil
}
