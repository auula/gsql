package gsql

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	_orm *ORM
	once sync.Once
)

type ORM struct {
	db        *sqlx.DB
	query     *Query
	parameter []interface{}
}

func Open(dbType string, databaseSourceName string) (err error) {
	var db *sqlx.DB
	once.Do(
		func() {
			if _orm != nil {
				_orm = new(ORM)
				db, err = sqlx.Open(dbType, databaseSourceName)
				_orm.db = db
			}
		},
	)
	return
}

func GetDB() *sqlx.DB {
	return _orm.db
}
func Close() {
	_orm.db.Close()
	_orm = nil
}
