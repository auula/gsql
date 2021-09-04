package gsql_test

import (
	"github.com/auula/gsql"
	"testing"
)

func TestSelectString(t *testing.T) {

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

	err, s := gsql.Select(UserInfo{}).From("user_info").Where("money >= ?", "100").Build()
	t.Log(err)
	t.Log(s)
}

func TestSelectMix(t *testing.T) {

	err, s := gsql.Select("name", 3.1415827, "age", 112, "money").
		From("user_info").
		Where("money >= ?", "100").Build()

	t.Log(err)
	t.Log(s)
}
