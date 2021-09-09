package gsql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	PrimaryKey     string
	Obj            interface{}
	IsSQLLimit     bool
	TableName      string
	DBType         int
	SelectColumns  *strings.Builder
	ConditionSQL   *strings.Builder
	OrderBySQL     *strings.Builder
	Parameter      []interface{}
	ExcludeColumns []string
}

type ActionResult struct {
	Err    error
	Result interface{}
}

type Action interface {
	Where
	Builder
	In() ActionResult
	One() ActionResult
	ById(id int) (err error)
	ByIds(ids ...int) ActionResult
	isNotNull()
	Exec() ActionResult
}

type Compare interface {
	GT() Where  // >
	Gte() Where // >=
	Lt() Where  // <
	Lte() Where // <=
	Equal() Where
	NotEq() Where
	Between() Where
}

// Builder generate structured query language code string
type Builder interface {
	Build() (error, string)
	String() string
	Buf() *strings.Builder
	Error(err error)
}

type Execution interface {
	SQLString(str string) Execution
	Execute() ActionResult
}

type Selector interface {
	Select() From
}

type From interface {
	From(model interface{}) Action
}

type Where interface {
	Where()
}

func Select() From {
	return &Query{
		IsSQLLimit:     false,
		TableName:      "",
		DBType:         0,
		SelectColumns:  new(strings.Builder),
		ConditionSQL:   new(strings.Builder),
		OrderBySQL:     new(strings.Builder),
		ExcludeColumns: make([]string, 0),
	}
}

func SelectAs(values []string) From {

	s := &Query{
		IsSQLLimit:     false,
		TableName:      "",
		DBType:         0,
		SelectColumns:  new(strings.Builder),
		ConditionSQL:   new(strings.Builder),
		OrderBySQL:     new(strings.Builder),
		ExcludeColumns: make([]string, 0),
	}

	for i, v := range values {
		s.SelectColumns.WriteString(v)
		if i == len(values)-1 {
			s.PrimaryKey = v
			break
		}
		s.SelectColumns.WriteString(", ")
	}

	return s
}

func (q *Query) From(model interface{}) Action {

	ty := reflect.TypeOf(model)
	q.TableName = ty.Name()

	if q.SelectColumns != nil && q.SelectColumns.String() == "" {
		for i := 0; i < ty.NumField(); i++ {

			if pkColumn := ty.Field(i).Tag.Get("pk"); pkColumn != "" {
				if q.PrimaryKey == "" {
					q.PrimaryKey = pkColumn
				}
			}

			q.SelectColumns.WriteString(ty.Field(i).Tag.Get("db"))
			if i == ty.NumField()-1 {
				break
			}
			q.SelectColumns.WriteString(", ")
		}
	}

	q.Obj = model

	return q
}

// 不能出现ById

func (q *Query) Where() {
	panic("implement me")
}

func (q *Query) In() ActionResult {
	panic("implement me")
}

func (q *Query) One() ActionResult {
	fmt.Println(q.Build())
	return ActionResult{Result: q.Obj}
}

func (q *Query) ById(id int) (err error) {

	var sql string
	if q.ConditionSQL != nil && q.ConditionSQL.String() == "" {
		q.ConditionSQL.WriteString(fmt.Sprintf(" %s = %d", q.PrimaryKey, id))
	}
	if err, sql = q.Build(); err != nil {
		return
	}

	return _db.Select(q.Obj, sql)

}

func (q *Query) ByIds(ids ...int) ActionResult {
	panic("implement me")
}

func (q *Query) isNotNull() {
	panic("implement me")
}

func (q *Query) Exec() ActionResult {
	panic("implement me")
}

func (q *Query) Build() (error, string) {

	sql := new(strings.Builder)
	sql.WriteString("SELECT ")

	if q.SelectColumns.String() != "" && q.SelectColumns != nil {
		sql.WriteString(q.SelectColumns.String())
	}

	if q.TableName == "" {
		return fmt.Errorf("%w", errors.New("table name found")), ""
	}

	sql.WriteString(fmt.Sprintf(" FROM %s", q.TableName))

	if q.ConditionSQL.String() != "" && q.ConditionSQL != nil {
		sql.WriteString(" WHERE")
		sql.WriteString(q.ConditionSQL.String())
	}

	return nil, sql.String()
}

func (q *Query) String() string {
	err, s := q.Build()
	if err != nil {
		return ""
	}
	return s
}

func (q *Query) Buf() *strings.Builder {
	panic("implement me")
}

func (q *Query) Error(err error) {
	panic("implement me")
}

func As(column string, asName string) string {
	return fmt.Sprintf("%s AS '%s'", column, asName)
}

func Alias(model interface{}, aliasMap map[string]string) []string {
	values, key := make([]string, 0), ""
	ty := reflect.TypeOf(model)
	for i := 0; i < ty.NumField(); i++ {

		if pkColumn := ty.Field(i).Tag.Get("pk"); pkColumn != "" {
			key = pkColumn
		}

		if v, ok := aliasMap[ty.Field(i).Tag.Get("db")]; ok {
			values = append(values, fmt.Sprintf("%s AS '%s'", ty.Field(i).Tag.Get("db"), v))
		} else {
			values = append(values, ty.Field(i).Tag.Get("db"))
		}
	}

	return append(values, key)
}
