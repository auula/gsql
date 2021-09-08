package gsql_test

import (
	"github.com/auula/gsql"
	"reflect"
	"testing"
)

type UserInfo struct {
	Name string `sql:"name"`
	Age  int    `sql:"age"`
}

func TestSelect(t *testing.T) {

	sql1 := gsql.Select().From(UserInfo{})
	t.Log(sql1)

	sql2 := gsql.SelectAs([]string{"name", gsql.As("age", "年龄")}).From(UserInfo{})
	t.Log(sql2)

	t.Log(reflect.DeepEqual(sql1, sql2))

	sql3 := gsql.SelectAs(gsql.Alias(UserInfo{}, map[string]string{
		"name": "名字",
	})).From(UserInfo{})

	t.Log(sql3)

	//=== RUN   TestSelect
	//sql_test.go:16: &{false UserInfo 0 name, age   []}
	//sql_test.go:19: &{false UserInfo 0 name, age AS '年龄'   []}
	//sql_test.go:21: false
	//sql_test.go:28: &{false UserInfo 0 name AS '名字', age   []}
	//--- PASS: TestSelect (0.00s)
	//PASS
}
