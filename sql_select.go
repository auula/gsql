package gsql

import (
	"github.com/auula/gsql/syntax"
	"reflect"
	"strings"
)

type SqlSelect struct {
	verb     map[string][]byte
	fields   []interface{}
	buf      *strings.Builder
	distinct bool
}

// Select sql.Select(user.id,user.name)
func Select(values ...interface{}) syntax.Selecter {

	s := &SqlSelect{
		verb: make(map[string][]byte, 10),
		buf:  new(strings.Builder),
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
func (sql *SqlSelect) Filter(filter syntax.Filter) syntax.Selecter {
	return sql
}

func (sql *SqlSelect) Distinct() syntax.Selecter {
	sql.distinct = true
	return sql
}

func (sql *SqlSelect) Where() syntax.Selecter {
	return sql
}

func (sql *SqlSelect) From() syntax.Selecter {
	return sql
}

func (sql *SqlSelect) Build() (error, string) {
	oldBuf := sql.buf.String()
	newBuf := new(strings.Builder)
	if sql.distinct {
		newBuf.WriteString("SELECT DISTINCT ")
		newBuf.WriteString(oldBuf)
		sql.buf = newBuf
		return nil, sql.buf.String()
	}
	newBuf.WriteString("SELECT ")
	newBuf.WriteString(oldBuf)
	sql.buf = newBuf
	return nil, sql.buf.String()
}
