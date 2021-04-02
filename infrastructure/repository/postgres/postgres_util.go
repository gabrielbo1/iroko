package postgres

import (
	"context"
	"database/sql"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/pkg"
	"github.com/jinzhu/now"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var ctx context.Context = context.Background()

//booleanToString - Convert bool type to postgres boolean string.
func booleanToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

//stringToString - Convert string type to postgres string.
func stringToString(s string) string {
	return "'" + s + "'"
}

//inCreate - Create sql instruction in with integer array.
func inCreate(ids []int) string {
	in := "("
	for i := 0; i < len(ids); i++ {
		in += strconv.Itoa(ids[i])
		in += ","
	}
	in = in[0 : len(ids)-1]
	return in + ")"
}

//datePostgres - Convert date string to sql type.
func datePostgres(dateString string) sql.NullTime {
	if dateString == "" {
		return sql.NullTime{}
	}

	var date time.Time
	var err error
	if date, err = now.Parse(dateString); err != nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  date,
		Valid: true,
	}
}

//generateSqlOffset - Generate offset instruction with page.
func generateSqlOffset(page pkg.Page) string {
	return "OFFSET " + strconv.Itoa((page.PageNumber)*page.PageSize) + " LIMIT " + strconv.Itoa(page.PageSize)
}

//prepareStmt - Prepare stmt.
func prepareStmt(ctx context.Context, tx *sql.Tx, repName, nameFunc, query string) (*sql.Stmt, *domain.Err) {
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, domain.NewErr().
			WithCode("postgres_util_10").
			WithMessage("Error prepare context, repository " + repName + " function " + nameFunc).
			WithError(err)
	}
	return stmt, nil
}

//scanStmt - Scan stmt.
func scanStmt(ctx context.Context, repName, nameFunc string, stmt *sql.Stmt, args ...interface{}) (bool, *domain.Err) {
	err := stmt.QueryRowContext(ctx).Scan(args)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		log.Println(err)
		return false, domain.NewErr().
			WithCode("postgres_util_20").
			WithMessage("Error scan stmt, respository " + repName + " function " + nameFunc).
			WithError(err)
	default:
		return true, nil
	}
	return true, nil
}

//scanParamStmt - Scan param stmt.
func scanParamStmt(nameRep, nameFunc string, stmt *sql.Stmt, query func(stmt *sql.Stmt) error) (bool, *domain.Err) {
	err := query(stmt)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		log.Println(err)
		return false, domain.NewErr().
			WithCode("postgres_util_30").
			WithMessage("Scan error, repository " + nameRep + " function " + nameFunc).
			WithError(err)
	default:
		return true, nil
	}
	return true, nil
}

//deleteNameColumn - Delete record with column name.
func deleteNameColumn(tx *sql.Tx, repName, table, nameCol string, id int) *domain.Err {
	sqlDelete := "DELETE FROM " + table + " WHERE " + nameCol + "=$1"
	stmt, err := prepareStmt(ctx, tx, repName, "Delete", sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt(repName, "Delete", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, id)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

//delete - Delete record with table name and id.
func delete(tx *sql.Tx, repName, table string, id int) *domain.Err {
	sqlDelete := "DELETE FROM " + table + " WHERE=$1"
	stmt, err := prepareStmt(ctx, tx, repName, "Delete", sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt(repName, "Delete", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, id)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}
