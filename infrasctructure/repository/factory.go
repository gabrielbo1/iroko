package repository

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/config"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/domain/permission"
	"github.com/gabrielbo1/iroko/infrasctructure/repository/postgres"
	log "github.com/sirupsen/logrus"
)

//DataBase - Defines the type of the constant to
// define supported databases.
type DataBase string

//PostgreSQL - PostgreSQL data base.
const PostgreSQL DataBase = "POSTGRESQL"

//DB - Pointer DB application.
var DB *sql.DB

//dataBase - Current data base application.
var dataBase *DataBase

//urlPostgresConnection - Url PostgreSQL driver
func urlPostgresConnection() string {
	connString := "host=" + config.EnvironmentVariableValue(config.BaseAddress)
	connString += " user=" + config.EnvironmentVariableValue(config.BaseUser)
	connString += " dbname=" + config.EnvironmentVariableValue(config.BaseName)
	connString += " password='" + config.EnvironmentVariableValue(config.BasePassword) + "'"
	connString += " sslmode=" + config.EnvironmentVariableValue(config.BaseSSL)
	return ""
}

//FindConnection - Find for connection to the database
// according to the Base parameter.
func FindConnection() (DB *sql.DB, err *domain.Err) {
	log.Info("FIND CONNECTION  DATA BASE.")
	var errDomain *domain.Err = nil
	if DB == nil {
		switch DataBase(config.EnvironmentVariableValue(config.Base)) {
		case PostgreSQL:
			*dataBase = PostgreSQL
			log.Info("Data base POSTGRESQL.")
			var err error
			DB, err = sql.Open("postgres", urlPostgresConnection())
			if err != nil {
				log.Info("FATAL ERROR CONNECTION POSTGRESQL.")
				log.Info(urlPostgresConnection())
				log.Fatal(err)
				errDomain = domain.NewErr().WithError(err)
			}
			break
		}
	}
	return DB, errDomain
}

//Transact - Transaction relational data bases like.
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

//NewCompanyRepository - Build domain.CompanyRepository with base configuration.
func NewCompanyRepository(tx *sql.Tx) domain.CompanyRepository {
	switch DataBase(*dataBase) {
	case PostgreSQL:
		return postgres.NewCompanyRepPostgres(tx)
	}
	return nil
}

//NewLoginRepository - Build domain.LoginRepository with base configuration.
func NewLoginRepository(tx *sql.Tx) domain.LoginRepository {
	switch DataBase(*dataBase) {
	case PostgreSQL:
		return postgres.NewLoginRepPostgres(tx)
	}
	return nil
}

//NewProfileRepository - Build permission.ProfileRepository with base configuration.
func NewProfileRepository(tx *sql.Tx) permission.ProfileRepository {
	switch DataBase(*dataBase) {
	case PostgreSQL:
		return postgres.NewProfileRepPostgres(tx)

	}
	return nil
}

//NewProfileCompanyRepository - Build permission.ProfileCompanyRepository with base configuration.
func NewProfileCompanyRepository(tx *sql.Tx) permission.ProfileCompanyRepository {
	switch DataBase(*dataBase) {
	case PostgreSQL:
		return postgres.NewProfileCompanyRepPostgres(tx)
	}
	return nil
}

//NewRoutineRepository - Build permission.RoutineRepository with base configuration.
func NewRoutineRepository(tx *sql.Tx) permission.RoutineRepository {
	switch DataBase(*dataBase) {
	case PostgreSQL:
		return postgres.NewRoutineRepPostgres(tx)
	}
	return nil
}
