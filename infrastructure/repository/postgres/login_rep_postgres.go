package postgres

import (
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
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
	sqlInsert := "INSERT INTO login(VERSION,LOGIN,PASSWORD,EMAIL) VALUES($1,$2,$3) RETURNING ID"
	stmt, err := prepareStmt(ctx, rep.tx, "loginRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("loginRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, 0, &login.Login, &login.Password, &login.Email).Scan(&id)
	}); !ok {
		err = errDomain
	}
	return
}

//Update - Update password login.
func (rep *loginRepPostgres) Update(login domain.Login) *domain.Err {
	sqlUpdate := "UPDATE login SET VERSION=$1, LOGIN=$2, PASSWORD=$3, EMAIL=$4 WHERE ID=$5"
	stmt, err := prepareStmt(ctx, rep.tx, "loginRepPostgres", "Update", sqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("loginRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, login.Version+1, &login.Login, &login.Password, &login.Email)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

//FindByLogin - Find login with login.
func (rep *loginRepPostgres) FindByLogin(id int) (domain.Login, *domain.Err) {
	sqlQuery := "SELECT ID, VERSION, LOGIN, PASSWORD, EMAIL FROM login WHERE ID=$1"
	stmt, errDomain := prepareStmt(ctx, rep.tx, "loginRepPostgres", "FindByLogin", sqlQuery)
	if errDomain != nil {
		return domain.Login{}, errDomain
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return domain.Login{}, domain.NewErr().
			WithCode("loginRepPostgres_10").
			WithMessage("Login repository error find by login.").
			WithError(err)
	}
	logins, err := parseLogin(rep.tx, rows)
	if err != nil {
		return domain.Login{}, domain.NewErr().
			WithCode("loginRepPostgres_10").
			WithMessage("Login repository error find by login.").
			WithError(err)
	}
	return logins[0], nil
}

//Delete - Delete record login with ID.
func (rep *loginRepPostgres) Delete(id int) *domain.Err {
	err := delete(rep.tx, "loginRepPostgres", "login", id)
	if err != nil {
		return err
	}
	return nil
}

func parseLogin(tx *sql.Tx, rows *sql.Rows) ([]domain.Login, error) {
	defer rows.Close()
	var result []domain.Login
	for rows.Next() {
		lg := domain.Login{}
		if err := rows.Scan(&lg.ID,
			&lg.Version,
			&lg.Login,
			&lg.Password,
			&lg.Email); err != nil {
			return nil, err
		}
		result = append(result, lg)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}
