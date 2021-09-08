# GSQL

GSQL is a structured query language code builder for golang.

## Genreate SQL

```go
package gsql_test

import (
	"fmt"
	"github.com/auula/gsql"
	"github.com/auula/gsql/syntax"
	"testing"
	"time"
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
		Name string `sql:"name"`
		Age  int    `sql:"age"`
	}
	// SELECT name, age FROM user_info WHERE money >= 999.90 AND age = 18
	err, s := gsql.Select(UserInfo{}).From("user_info").Where("money >= ? AND age = ?", 999.9, 18).Build()
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

	// SELECT name, age, money AS '余额' FROM user_info WHERE name = 'Leon Ding' LIMIT 1
	sql := gsql.Select("name", "age", syntax.As("money", "余额")).
		From("user_info").
		Where("name = ?", "Leon Ding").Limit(1).String()

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
		Name  string  `sql:"name"`
		Age   int     `sql:"age"`
		Money float64 `sql:"money"`
	}

	sql := gsql.SelectAs(syntax.Alias(UserInfo{}, map[string]string{
		"name":  "用户名",
		"money": "金钱",
	})).
		From("user_info").
		Where("name = ?", "Leon Ding").Limit(1).String()
	t.Log(sql)

	syntaxSql := gsql.Select(UserInfo{}).From("user_info").Limit(2)

	err, s := syntax.Limit(syntaxSql, true, 1, 3).Build()

	t.Log(err)
	t.Log(s)

	//=== RUN   TestSelectAlias
	//sql_test.go:92: SELECT name AS '用户名', age, money AS '金钱' FROM user_info WHERE name = 'Leon Ding' LIMIT 1
	//sql_test.go:98: limit syntax recurring
	//sql_test.go:99: SELECT name, age, money FROM user_info LIMIT 2
	//--- PASS: TestSelectAlias (0.00s)
	//PASS

}

func TestSqlSelectOrderBy(t *testing.T) {

	type UserInfo struct {
		Name  string  `sql:"name"`
		Age   int     `sql:"age"`
		Money float64 `sql:"money"`
	}

	syntaxSql := gsql.SelectAs(syntax.Alias(UserInfo{}, map[string]string{
		"name": "用户名",
	})).From("user_info")

	// SELECT name AS '用户名', age, money FROM user_info ORDER BY  money DESC, age ASC
	sql := syntax.OrderBy(syntaxSql, []syntax.OrderRow{
		{"money", syntax.DESC},
		{"age", syntax.ASC},
	})

	s := syntax.Limit(sql, true, 1, 3).String()

	t.Log(s)
}

func TestCol(t *testing.T) {
	// name = 'Leon Ding' AND age = 19
	sql := syntax.Condition(syntax.Col("name").Equal("'Leon Ding'")).
		AND(syntax.Col("age").Equal(19))
	t.Log(sql)
}

func TestIN(t *testing.T) {
	// name IN ('Jaco','Kimi')
	sql := syntax.Condition(syntax.Col("name").In([]string{"Jaco", "Kimi"})).String()
	t.Log(sql)
}

func TestLike(t *testing.T) {
	// name LIKE '%Di%'
	sql := syntax.Condition(syntax.Col("name").Like("%Di%")).String()
	t.Log(sql)
}

func TestBetween(t *testing.T) {

	// SELECT name, age, money AS '余额' FROM user_info
	// WHERE created_at BETWEEN '2000-01-08 00:00:00'
	// AND '2021-09-07 20:38:52'
	// AND age BETWEEN 10 AND 21

	err, left := syntax.Col("created_at").Between([]interface{}{
		"'2000-01-08 00:00:00'",
		fmt.Sprintf("'%v'", time.Now().Format("2006-01-02 15:04:05")),
	})

	err, right := syntax.Col("age").Between([]interface{}{10, 21})
	sql := syntax.Condition(left).AND(right)

	syntaxSql := gsql.Select("name", "age", syntax.As("money", "余额")).
		From("user_info").WhereBind(sql).String()

	if err != nil {
		t.Error(err)
	}

	t.Log(syntaxSql)
}
```
