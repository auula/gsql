package gsql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Inserter interface {
	Into
	Executor
}

type Into interface {
	Values(v ...interface{}) Builder
}

type Executor interface {
	Exec() Builder
}

type Execute struct {
	TableName string
	Columns   []string
	Value     []string
	Obj       interface{}
	Err       error
}

func Insert(model interface{}, filter []string) Into {

	e := &Execute{
		TableName: "",
		Columns:   make([]string, 0, 20),
		Value:     make([]string, 0, 20),
		Obj:       model,
	}

	typeOf := reflect.TypeOf(model)
	e.TableName = typeOf.Name()

	for i := 0; i < typeOf.NumField(); i++ {
		e.Columns = append(e.Columns, typeOf.Field(i).Tag.Get("db"))
	}

	for i := 0; i < len(e.Columns); i++ {
		for _, v := range filter {
			if e.Columns[i] == v {
				e.Columns = append(e.Columns[:i], e.Columns[i+1:]...)
			}
		}
	}

	return e
}

func (e *Execute) Values(v ...interface{}) Builder {
	if len(v) != len(e.Columns) {
		e.Err = fmt.Errorf("missing parameters: %w", errors.New("where syntax lack of conditions"))
		return e
	}
	for _, value := range v {
		switch value.(type) {
		case string:
			e.Value = append(e.Value, fmt.Sprintf("'%s'", value.(string)))
		case float64:
			e.Value = append(e.Value, fmt.Sprintf("%v", value))
		case int:
			e.Value = append(e.Value, fmt.Sprintf("%v", value))
		default:
			continue
		}
	}
	return e
}

func (e *Execute) Build() (error, string) {

	sql := new(strings.Builder)
	sql.WriteString("INSERT INTO ")
	sql.WriteString(e.TableName)

	columnsSql := new(strings.Builder)

	if len(e.Columns) > 0 {

		for i, column := range e.Columns {
			columnsSql.WriteString(fmt.Sprintf("%s", column))
			if i == len(e.Columns)-1 {
				break
			}
			columnsSql.WriteString(", ")
		}

		sql.WriteString(fmt.Sprintf(" (%v)", columnsSql))

	}

	sql.WriteString(" VALUES ")

	valueSql := new(strings.Builder)

	if len(e.Value) > 0 {

		for i, v := range e.Value {
			valueSql.WriteString(fmt.Sprintf("%s", v))
			if i == len(e.Value)-1 {
				break
			}
			valueSql.WriteString(", ")
		}

		sql.WriteString(fmt.Sprintf("(%v)", valueSql))

	}

	return e.Err, sql.String()
}

func (e *Execute) String() string {
	err, s := e.Build()
	if err != nil {
		return ""
	}
	return s
}
