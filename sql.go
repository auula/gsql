package gsql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/auula/gsql/syntax"
)

type SqlSelect struct {
	buf      *strings.Builder
	distinct bool
	Err      error
}

func (sql *SqlSelect) WhereBind(columned syntax.Columned) syntax.Builder {
	sql.buf.WriteString(" WHERE ")
	sql.buf.WriteString(columned.String())
	return sql
}

func (sql *SqlSelect) Limit(row int) syntax.Filter {
	sql.buf.WriteString(fmt.Sprintf(" LIMIT %v", row))
	return sql
}

func (sql *SqlSelect) OrderBy(row []syntax.OrderRow) syntax.Filter {
	syntax.OrderBy(sql, row)
	return sql
}

// Select sql.Selector(user.id,user.name)
func Select(values ...interface{}) syntax.Form {

	s := &SqlSelect{
		buf: new(strings.Builder),
	}

	if len(values) == 1 {
		ty := reflect.TypeOf(values[0])
		for i := 0; i < ty.NumField(); i++ {
			s.buf.WriteString(ty.Field(i).Tag.Get("sql"))
			if i == ty.NumField()-1 {
				break
			}
			s.buf.WriteString(", ")
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

func SelectAs(values []string) syntax.Form {

	s := &SqlSelect{
		buf: new(strings.Builder),
	}

	for i, v := range values {

		s.buf.WriteString(v)
		if i == len(values)-1 {
			break
		}
		s.buf.WriteString(", ")

	}
	return s
}

func (sql *SqlSelect) Buf() *strings.Builder {
	return sql.buf
}

func (sql *SqlSelect) Distinct() syntax.Selector {
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

func (sql *SqlSelect) From(tab string) syntax.Filter {
	sql.buf.WriteString(" FROM ")
	sql.buf.WriteString(tab)
	return sql
}

func (sql *SqlSelect) String() string {
	_, s := sql.Build()
	if sql.Err != nil {
		return ""
	}
	return s
}

func (sql *SqlSelect) Error(err error) {
	sql.Err = fmt.Errorf("%w", err)
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
	return sql.Err, sql.buf.String()
}
