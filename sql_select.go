package gsql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/auula/gsql/syntax"
)

type SqlSelect struct {
	tableName string
	buf       *strings.Builder
	distinct  bool
	Err       error
}

func (sql *SqlSelect) Limit(x int, y int) syntax.Filter {
	panic("implement me")
}

func (sql *SqlSelect) Order(field interface{}, sort syntax.SortType) syntax.Filter {
	panic("implement me")
}

// Select sql.Select(user.id,user.name)
func Select(values ...interface{}) syntax.Form {

	s := &SqlSelect{
		buf: new(strings.Builder),
	}

	if len(values) == 1 {
		ty := reflect.TypeOf(values[0])
		for i := 0; i < ty.NumField(); i++ {
			s.buf.WriteString(ty.Field(i).Tag.Get("json"))
			if i == ty.NumField()-1 {
				break
			}
			s.buf.WriteString(", ")
		}
		if s.tableName == "" {
			s.tableName = ty.Name()
		}
		return s
	}

	for i, v := range values {
		switch v.(type) {
		case string:
			s.buf.WriteString(v.(string))
			if i == len(values)-1 {
				break
			}
			s.buf.WriteString(", ")
		default:
			continue
		}
	}
	return s
}
func (sql *SqlSelect) Buf() *strings.Builder {
	return sql.buf
}

func (sql *SqlSelect) Distinct() syntax.Select {
	sql.distinct = true
	return nil
}

// Where money >= 100 "money > #?" 100
func (sql *SqlSelect) Where(s string, v ...interface{}) syntax.Filter {

	buf := new(strings.Builder)
	if len(v) != strings.Count(s, "?") {
		sql.Err = fmt.Errorf("missing parameters: %w", errors.New("where syntax lack of conditions"))
		return sql
	}
	buf.WriteString(" WHERE ")
	for _, value := range v {
		switch value.(type) {
		case string:
			buf.WriteString(strings.Replace(s, "?", fmt.Sprintf("'%s'", value.(string)), -1))
			// 数字类型 时间类型 浮点数类型
		default:
			continue
		}
	}
	sql.buf.WriteString(buf.String())
	return sql
}

func (sql *SqlSelect) From(tab string) syntax.Select {
	if sql.tableName == "" || tab != "" {
		sql.tableName = tab
	}
	return sql
}

func (sql *SqlSelect) Build() (error, string) {
	oldBuf := sql.buf.String()
	newBuf := new(strings.Builder)
	if sql.distinct {
		newBuf.WriteString("SELECT DISTINCT ")
		newBuf.WriteString(oldBuf)
		sql.buf = newBuf
	}
	newBuf.WriteString("SELECT ")
	newBuf.WriteString(oldBuf)
	sql.buf = newBuf
	sql.buf.WriteString(" FROM ")
	sql.buf.WriteString(sql.tableName)
	return sql.Err, sql.buf.String()
}
