package gsql_test

import (
	"github.com/auula/gsql"
	"testing"
)

func TestInsert(t *testing.T) {
	// INSERT INTO UserInfo (id, name, age) VALUES (1001, 'Tom', 21)
	sql := gsql.Insert(UserInfo{}, nil).Values(1001, "Tom", 21).String()
	t.Log(sql)
}

func TestInsertFilter(t *testing.T) {
	// INSERT INTO UserInfo (name, age) VALUES ('Tom', 21)
	sql := gsql.Insert(UserInfo{}, []string{"id"}).Values("Tom", 21).String()
	t.Log(sql)
}
