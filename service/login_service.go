package service

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/infrastructure/repository"
)

//CreateLogin - Create a new login.
func CreateLogin(login *domain.Login) (err *domain.Err) {
	if _, err = login.ValidLogin(); err != nil {
		return
	}
	repository.Transact(repository.DB, func(tx *sql.Tx) error {
		_, err = repository.NewLoginRepository(tx).Save(*login)
		return nil
	})
	return
}

//UpdateLogin - Update a login record.
func UpdateLogin(login *domain.Login) (err *domain.Err) {
	if _, err = login.ValidLogin(); err != nil {
		return
	}
	repository.Transact(repository.DB, func(tx *sql.Tx) error {
		err = repository.NewLoginRepository(tx).Update(*login)
		return nil
	})
	return
}

//DeleteLogin - Delete login record.
func DeleteLogin(id int) (err *domain.Err) {
	repository.Transact(repository.DB, func(tx *sql.Tx) error {
		err = repository.NewLoginRepository(tx).Delete(id)
		return nil
	})
	return
}
