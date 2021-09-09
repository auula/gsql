package gsql_test

import (
	"github.com/auula/gsql"
	"reflect"
	"testing"
)

type UserInfo struct {
	Id   string `db:"id" pk:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func TestSelect(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE id = 1
	sql1 := gsql.Select().From(UserInfo{}).ById(1)
	t.Log(sql1)

	sql2 := gsql.SelectAs([]string{"name", gsql.As("age", "年龄")}).From(UserInfo{}).ById(2)
	t.Log(sql2)

	t.Log(reflect.DeepEqual(sql1, sql2))

	sql3 := gsql.SelectAs(gsql.Alias(UserInfo{}, map[string]string{
		"name": "名字",
	})).From(UserInfo{}).ById(1)

	t.Log(sql3)

	//=== RUN   TestSelect
	//sql_test.go:16: &{false UserInfo 0 name, age   []}
	//sql_test.go:19: &{false UserInfo 0 name, age AS '年龄'   []}
	//sql_test.go:21: false
	//sql_test.go:28: &{false UserInfo 0 name AS '名字', age   []}
	//--- PASS: TestSelect (0.00s)
	//PASS
}
