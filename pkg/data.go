package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// InCreate - Format int array to SQL in instruction.
func InCreate(ids []int) string {
	if len(ids) > 0 {
		idsStr := make([]string, len(ids))
		for i, id := range ids {
			idsStr[i] = strconv.Itoa(id)
		}
		return "(" + strings.Join(idsStr, ",") + ")"
	}
	return ""
}

// PeprareStmt - Prepare sql stmt  with context.
func PeprareStmt(tx *sql.Tx, ctx context.Context, repositoryName, funcName, query string) (*sql.Stmt, *Err) {
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		errMsg := fmt.Sprintf("DATA_10: Error prepare stmt, repository %s, function %s", repositoryName, funcName)
		log.Println(errMsg)
		return nil, NewErr().WithCode("DATA_10").WithMessage(errMsg)
	}
	return stmt, nil
}

func ScanStmt(repositoryName, funcName string, stmt *sql.Stmt, ctx context.Context, args ...interface{}) (bool, *Err) {
	err := stmt.QueryRowContext(ctx).Scan(args)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		errMsg := fmt.Sprintf("DATA_20: Error scan stmt, repository %s, function %s", repositoryName, funcName)
		return false, NewErr().WithCode("DATA_20").WithMessage(errMsg)
	}
	return true, nil
}

// ScanParamStmt - Scan parameters in query statement.
func ScanParamStmt(repositoryName, funcName string, stmt *sql.Stmt, query func(stmt *sql.Stmt) error) (bool, *Err) {
	err := query(stmt)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		errMsg := fmt.Sprintf("DATA_20: Error scan param, repository %s, function %s", repositoryName, funcName)
		return false, NewErr().WithCode("DATA_20").WithMessage(errMsg)
	}
	return true, nil
}

// Page - Paginated queries pattern.
type Page struct {

	//First - Is first page.
	First bool `json:"first"`

	//Last - Is last page.
	Last bool `json:"last"`

	//PageNumber - The page number.
	PageNumber int `json:"pageNumber"`

	//PageSize - The page size.
	PageSize int `json:"pageSize"`

	//Content - The content page.
	Content interface{} `json:"content"`
}

// Repository - Define basic operations entities.
type Repository interface {
	Count(tx *sql.Tx) (int, *Err)

	Delete(tx *sql.Tx, entidade interface{}) *Err

	DeleteAll(tx *sql.Tx) *Err

	DeleteSlice(tx *sql.Tx, entidades []interface{}) *Err

	DeleteById(tx *sql.Tx, id string) *Err

	ExistsById(tx *sql.Tx, id string) (bool, *Err)

	FindAll(tx *sql.Tx) (entidades []interface{}, erro *Err)

	FindAllById(tx *sql.Tx, ids []string) (entidades []interface{}, erro *Err)

	FindById(tx *sql.Tx, id string) (entidades interface{}, erro *Err)

	Save(tx *sql.Tx, entidade interface{}) (string, *Err)

	SaveAll(tx *sql.Tx, entidades []interface{}) *Err

	Update(tx *sql.Tx, entidade interface{}) *Err

	PaginatedQuery(tx *sql.Tx, page Page, fieldsFilter map[string]interface{}) (Page, *Err)
}

const (
	POSTGRESQL string = "POSTGRESQL"
)

// SGDB connection pointer.
var db *sql.DB

func DB() (*sql.DB, *Err) {
	if ConfigVars.EnvironmentVariableValue(Base) == POSTGRESQL && db == nil {
		var err error
		connStr := fmt.Sprintf("host=%s user=%s dbname=%s password='%s' sslmode=disable",
			ConfigVars.EnvironmentVariableValue(BaseAddress),
			ConfigVars.EnvironmentVariableValue(BaseUser),
			ConfigVars.EnvironmentVariableValue(BaseName),
			ConfigVars.EnvironmentVariableValue(BasePassword))
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
			return nil, NewErr().WithCode("DATA_30").
				WithMessagef("Database connection error, url %s", connStr)
		}
	}
	return db, nil
}
