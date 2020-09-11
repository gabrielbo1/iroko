package repository

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/config"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/domain/permission"
	"github.com/gabrielbo1/iroko/infrastructure/repository/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	migPostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"strings"
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

func init() {
	var k DataBase = "NO_DATABASE"
	dataBase = &k
}

func getDataBase() *DataBase {
	if dataBase == nil {
		switch config.EnvironmentVariableValue(config.Base) {
		case string(PostgreSQL):
			var dt DataBase = PostgreSQL
			dataBase = &dt
			break
		default:
			log.Fatal("Invalid data base configuration.")
		}
	}
	return dataBase
}

//urlPostgresConnection - Url PostgreSQL driver
func urlPostgresConnection() string {
	connString := "host=" + config.EnvironmentVariableValue(config.BaseAddress)
	connString += " user=" + config.EnvironmentVariableValue(config.BaseUser)
	connString += " dbname=" + config.EnvironmentVariableValue(config.BaseName)
	connString += " password='" + config.EnvironmentVariableValue(config.BasePassword) + "'"
	connString += " sslmode=" + config.EnvironmentVariableValue(config.BaseSSL)
	return connString
}

//MigrationInit - Migration execute.
func MigrationInit() {
	var migrator *migrate.Migrate
	var err error
	var errDomain *domain.Err
	var sourceDriver source.Driver
	var dbInstance database.Driver

	if DB, errDomain = FindConnection(); errDomain != nil {
		log.Fatal(errDomain)
	}

	switch DataBase(*getDataBase()) {
	case PostgreSQL:
		f := &file.File{}
		if sourceDriver, err = f.Open("file://" + "./infrastructure/repository/postgres/migration"); err != nil {
			log.Fatal(err)
		}

		if dbInstance, err = migPostgres.WithInstance(DB, &migPostgres.Config{}); err != nil {
			log.Fatal(err)
		}

		if migrator, err = migrate.NewWithInstance(
			"file",
			sourceDriver,
			config.EnvironmentVariableValue(config.BaseName),
			dbInstance); err != nil {
			log.Fatal(err)
		}

		if err = migrator.Up(); err != nil && !strings.Contains(err.Error(), "no change") {
			log.Fatal(err)
		}
		break
	}
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
