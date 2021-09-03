package gsql_test

import (
	"github.com/auula/gsql"
	"testing"
)

func TestSelectString(t *testing.T) {

	s := gsql.Select("name", "age", "money")

	t.Log(s.(*gsql.SqlSelect).Buf())
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