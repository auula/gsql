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

	// SELECT id, name, age FROM UserInfo
	sql1 := gsql.Select().From(UserInfo{})
	t.Log(sql1)

	// SELECT name, age AS '年龄' FROM UserInfo WHERE  id = 2
	sql2 := gsql.SelectAs([]string{"name", gsql.As("age", "年龄"), "id"}).From(UserInfo{}).ById(2)
	t.Log(sql2)

	// SELECT id, name AS '名字', age FROM UserInfo WHERE  id = 1
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

func TestSelectIns(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE age IN (21, 19, 28)
	sql := gsql.Select().From(UserInfo{}).In("age", 21, 19, 28)
	t.Log(sql)

	sql2 := gsql.Select().From(UserInfo{}).In("name", []string{
		"Jaco",
		"Leon",
	})
	t.Log(sql2)
}

func TestSelectOne(t *testing.T) {

	// SELECT id, name, age FROM UserInfo LIMIT 1
	_, sql := gsql.Select().From(UserInfo{}).One()
	t.Log(sql)

	// SELECT id, name, age FROM UserInfo WHERE age > 10 LIMIT 1
	err, sql2 := gsql.Select().From(UserInfo{}).Where("age > ?", 10).One()

	if err == nil {
		t.Log(sql2)
	}

}

func TestSelectLimit(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE age > 10 LIMIT 3 OFFSET 1
	sql2 := gsql.Select().From(UserInfo{}).Where("age > ?", 10).Limit(true, 1, 3)

	t.Log(sql2)

}

func TestSelectOrder(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE age > 10 ORDER BY id ASC LIMIT 3 OFFSET 1
	sql2 := gsql.Select().From(UserInfo{}).Where("age > ?", 10).Order([]gsql.Rows{
		{"id", "ASC"},
	}).Limit(true, 1, 3)

	t.Log(sql2)

}
