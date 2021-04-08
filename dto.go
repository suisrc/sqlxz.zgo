package sqlxz

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var (
	TBL = "zgo_"
)

type EntityDB interface {
	Get(dest interface{}, query string, args ...interface{}) error
}
type ArrayDB interface {
	Select(dest interface{}, query string, args ...interface{}) error
}

type ExecDB interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

var _ EntityDB = &sqlx.DB{}
var _ ArrayDB = &sqlx.DB{}
var _ ExecDB = &sqlx.DB{}
