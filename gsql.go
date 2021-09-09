package gsql

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	_db  *sqlx.DB
	once sync.Once
)

func Open(dbType string, databaseSourceName string) (err error) {
	once.Do(
		func() {
			if _db != nil {
				_db, err = sqlx.Open(dbType, databaseSourceName)
			}
		},
	)
	return
}

func GetDB() *sqlx.DB {
	if _db != nil {
		return _db
	}
	return nil
}

func Close() {
	if _db != nil {
		_db.Close()
		_db = nil
	}
}
