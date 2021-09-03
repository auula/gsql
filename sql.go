package gsql

import (
	"github.com/auula/gsql/verb"
)

var sqlBuilder *sql

func init() {
	sqlBuilder = new(sql)
}

type sql struct {
}

func (sql *sql) Select() verb.Selecter {
	return sql
}

func (sql *sql) Filter(filter verb.Filter) verb.Selecter {
	return sql
}

func (sql *sql) Distinct() verb.Selecter {
	return sql
}

func (sql *sql) Where() verb.Selecter {
	return sql
}

func (sql *sql) From() verb.Selecter {
	return sql
}

func (sql *sql) Build() (error, string) {
	return nil, ""
}
