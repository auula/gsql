package gsql_test

import (
	"github.com/auula/gsql"
	"testing"
)

func TestSelectString(t *testing.T) {

	err, s := gsql.Select("name", "age", "money").Distinct().Build()
	if err != nil {
		t.Log(err)
	}

	t.Log(s)
}

func TestSelectTag(t *testing.T) {

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	s := gsql.Select(UserInfo{})

	t.Log(s.(*gsql.SqlSelect).Buf())
}

func TestSelectMix(t *testing.T) {

	s := gsql.Select("name", 3.1415827, "age", 112, "money")

	t.Log(s.(*gsql.SqlSelect).Buf())
}
