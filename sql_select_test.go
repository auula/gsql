package gsql_test

import (
	"github.com/auula/gsql"
	"testing"
)

func TestSelectString(t *testing.T) {
	// SELECT name, age, money FROM user_info WHERE money >= '100'
	err, s := gsql.Select("name", "age", "money").
		From("user_info").Where("money >= ?", "100").
		Build()

	t.Log(err)
	t.Log(s)
}

func TestSelectTag(t *testing.T) {

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	// SELECT name, age FROM user_info WHERE money >= '100'
	err, s := gsql.Select(UserInfo{}).From("user_info").Where("money >= ?", "100").Build()
	t.Log(err)
	t.Log(s)
}

func TestSelectMix(t *testing.T) {

	//sql_select_test.go:36: missing parameters: where syntax lack of conditions
	//sql_select_test.go:37: SELECT name, age, money FROM user_info
	err, s := gsql.Select("name", 3.1415827, "age", 112, "money").
		From("user_info").
		Where("money >= ?", "100", "1").Build()

	t.Log(err)
	t.Log(s)
}

func TestSelectToStr(t *testing.T) {

	// SELECT name, age, money FROM user_info WHERE name = 'Leon Ding'
	sql := gsql.Select("name", "age", "money").
		From("user_info").
		Where("name = ?", "Leon Ding").String()

	t.Log(sql)
}
