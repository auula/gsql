package gsql_test

import (
	"testing"

	"github.com/auula/gsql"
	"github.com/auula/gsql/syntax"
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

	//sql_test.go:36: missing parameters: where syntax lack of conditions
	//sql_test.go:37: SELECT name, age, money FROM user_info
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

func TestSelectAsName(t *testing.T) {

	// SELECT name, age, money AS '余额' FROM user_info WHERE name = 'Leon Ding'
	sql := gsql.Select("name", "age", syntax.As("money", "余额")).
		From("user_info").
		Where("name = ?", "Leon Ding").String()

	t.Log(sql)
}

func TestSelectFilter(t *testing.T) {

	// offset=true SELECT name, age, money AS '余额' FROM user_info LIMIT 3 OFFSET 1
	// offset=false SELECT name, age, money AS '余额' FROM user_info LIMIT 1,3

	syntaxSql := gsql.Select("name", "age", syntax.As("money", "余额")).
		From("user_info")

	// SELECT name, age, money AS '余额' FROM user_info LIMIT 3 OFFSET 1
	sql := syntax.Limit(syntaxSql, true, 1, 3).String()
	t.Log(sql)

}

func TestSelectAlias(t *testing.T) {

	type UserInfo struct {
		Name  string  `json:"name"`
		Age   int     `json:"age"`
		Money float64 `json:"money"`
	}

	sql := gsql.SelectAs(syntax.Alias(UserInfo{}, map[string]string{
		"name":  "用户名",
		"money": "金钱",
	})).
		From("user_info").
		Where("name = ?", "Leon Ding")

	code := syntax.Limit(sql, true, 1, 3).String()
	t.Log(code)

	syntaxSql := gsql.SelectAs(syntax.Alias(UserInfo{}, map[string]string{
		"name":  "用户名",
		"money": "金钱",
	})).
		From("user_info")

	err, s := syntax.Limit(syntaxSql, true, 1, 3).Build()

	t.Log(err)
	t.Log(s)

	//=== RUN   TestSelectAlias
	//sql_test.go:94: SELECT name AS '用户名', age, money AS '金钱' FROM user_info WHERE name = 'Leon Ding' LIMIT 3 OFFSET 1
	//sql_test.go:104: limit syntax recurring
	//sql_test.go:105: SELECT name AS '用户名', age, money AS '金钱' FROM user_info LIMIT 1 OFFSET 1
	//--- PASS: TestSelectAlias (0.00s)
	//PASS

}

func TestSqlSelectOrderBy(t *testing.T) {
	type UserInfo struct {
		Name  string  `json:"name"`
		Age   int     `json:"age"`
		Money float64 `json:"money"`
	}

	syntaxSql := gsql.SelectAs(syntax.Alias(UserInfo{}, map[string]string{
		"name": "用户名",
	})).From("user_info")

	// SELECT name AS '用户名', age, money FROM user_info ORDER BY  money DESC, age ASC
	sql := syntax.OrderBy(syntaxSql, []syntax.OrderRow{
		{"money", syntax.DESC},
		{"age", syntax.ASC},
	}).String()
	t.Log(sql)
}
