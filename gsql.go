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
	once.Do(
		func() {
			if _orm != nil {
				_orm = new(ORM)
				_orm.db, err = sqlx.Open(dbType, databaseSourceName)
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
	if _orm != nil {
		_orm = nil
	}
}
