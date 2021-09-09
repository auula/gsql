package gsql_test

import (
	"github.com/auula/gsql"
	"testing"
)

type UserInfo struct {
	Id   int    `db:"id" pk:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func TestSelect(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE id = 1
	sql1 := gsql.Select().From(UserInfo{}).ById(1)
	t.Log(sql1)

	// SELECT id, name AS '名字', age, id FROM UserInfo WHERE id = 1
	sql2 := gsql.SelectAs([]string{"name", gsql.As("age", "年龄"), "id"}).From(UserInfo{}).ById(2)
	t.Log(sql2)

	// SELECT id, name AS '名字', age, id FROM UserInfo WHERE id = 1
	sql3 := gsql.SelectAs(gsql.Alias(UserInfo{}, map[string]string{
		"name": "名字",
	})).From(UserInfo{}).ById(1)
	t.Log(sql3)

}

func TestSelectByIds(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE id IN (1, 2, 3)
	sql := gsql.Select().From(UserInfo{}).ByIds(1, 2, 3)
	t.Log(sql)

}
