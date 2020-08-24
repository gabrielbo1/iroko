package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/domain/permission"
)

type routineRepPostgres struct {
	tx *sql.Tx
}

//NewRoutineRepPostgres - Implementation permission. with PostgreSQL.
func NewRoutineRepPostgres(tx *sql.Tx) *routineRepPostgres {
	return &routineRepPostgres{tx: tx}
}

//Save - Save new Routine entity.
func (rep *routineRepPostgres) Save(profile permission.Routine) (int, *domain.Err) {
	return 0, nil
}

//Update - Update Routine entity.
func (rep *routineRepPostgres) Update(profile permission.Routine) *domain.Err {
	return nil
}

//FindById - Find Routine with id.
func (rep *routineRepPostgres) FindById(id int) (permission.Routine, *domain.Err) {
	return permission.Routine{}, nil
}

//FindAll - Find all Routine entities.
func (rep *routineRepPostgres) FindAll() ([]permission.Routine, *domain.Err) {
	return []permission.Routine{}, nil
}
