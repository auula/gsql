package gsql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	PrimaryKey     string
	TableName      string
	DBType         int
	SelectColumns  *strings.Builder
	ConditionSQL   *strings.Builder
	OrderBySQL     *strings.Builder
	SQLLimit       *strings.Builder
	ExcludeColumns []string
	Err            error
}

type Rows struct {
	Field string
	Sort  string
}

type Action interface {
	Where
	Limit
	Builder
	In(column string, values ...interface{}) Builder
	One() (error, string)
	ById(id int) Builder
	ByIds(ids ...int) Builder
}

type Limit interface {
	Limit(offset bool, index, row int) Builder
}

type Order interface {
	Order(rows []Rows) Action
}

// Builder generate structured query language code string
type Builder interface {
	Build() (error, string)
	String() string
}

type Selector interface {
	Select() From
}

type From interface {
	From(model interface{}) Action
}

type Where interface {
	Where(sql string, values ...interface{}) Action
}

func Select() From {
	return &Query{
		TableName:      "",
		DBType:         0,
		SQLLimit:       new(strings.Builder),
		SelectColumns:  new(strings.Builder),
		ConditionSQL:   new(strings.Builder),
		OrderBySQL:     new(strings.Builder),
		ExcludeColumns: make([]string, 0),
	}
}

func SelectAs(values []string) From {

	s := &Query{
		TableName:      "",
		DBType:         0,
		SQLLimit:       new(strings.Builder),
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

func (q *Query) Limit(offset bool, index, row int) Builder {

	if q.SQLLimit != nil {

		if offset {
			q.SQLLimit.WriteString(fmt.Sprintf("%v OFFSET %v", row, index))
			return q
		}

		q.SQLLimit.WriteString(fmt.Sprintf("%v,%v", index, row))
	}

	return q
}

func (q *Query) Order(rows []Rows) Action {
	panic("implement me")
}

func (q *Query) From(model interface{}) Action {
	ty := reflect.TypeOf(model)
	q.TableName = ty.Name()
	if q.SelectColumns != nil {
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

	return q
}

// 不能出现ById

func (q *Query) Where(sql string, values ...interface{}) Action {
	if len(values) != strings.Count(sql, "?") {
		q.Err = fmt.Errorf("missing parameters: %w", errors.New("where syntax lack of conditions"))
		return q
	}
	for _, value := range values {
		switch value.(type) {
		case string:
			sql = strings.Replace(sql, "?", fmt.Sprintf("'%s'", value.(string)), 1)
			// 数字类型 时间类型 浮点数类型
		case float64:
			sql = strings.Replace(sql, "?", fmt.Sprintf("%.2f", value.(float64)), 1)
		case int:
			sql = strings.Replace(sql, "?", fmt.Sprintf("%d", value.(int)), 1)
		default:
			continue
		}
	}
	q.ConditionSQL.WriteString(sql)
	return q
}

func (q *Query) In(column string, values ...interface{}) Builder {
	buf := new(strings.Builder)
	for i, v := range values {
		switch v.(type) {
		case []string:
			for i, s := range v.([]string) {
				buf.WriteString(fmt.Sprintf("'%s'", s))
				if i == len(v.([]string))-1 {
					break
				}
				buf.WriteString(", ")
			}
		case []int:
			for i, s := range v.([]int) {
				buf.WriteString(fmt.Sprintf("'%d'", s))
				if i == len(v.([]int))-1 {
					break
				}
				buf.WriteString(", ")
			}
		default:
			buf.WriteString(fmt.Sprintf("%v", v))
		}

		if i == len(values)-1 {
			break
		}
		buf.WriteString(", ")
	}
	if q.ConditionSQL != nil {
		q.ConditionSQL.WriteString(fmt.Sprintf(" %s IN (%v)", column, buf))
	}

	return q
}

func (q *Query) One() (error, string) {
	err, s := q.Build()
	if err != nil {
		return err, ""
	}
	return nil, fmt.Sprintf("%s LIMIT 1", s)
}

func (q *Query) ById(id int) Builder {

	if q.ConditionSQL != nil {
		q.ConditionSQL.WriteString(fmt.Sprintf(" %s = %d", q.PrimaryKey, id))
	}

	return q

}

func (q *Query) ByIds(ids ...int) Builder {

	buf := new(strings.Builder)
	for i, id := range ids {
		buf.WriteString(fmt.Sprintf("%d", id))
		if i == len(ids)-1 {
			break
		}
		buf.WriteString(", ")
	}
	if q.ConditionSQL != nil {
		q.ConditionSQL.WriteString(fmt.Sprintf(" %s IN (%v)", q.PrimaryKey, buf))
	}

	return q
}

func (q *Query) Build() (error, string) {

	sql := new(strings.Builder)
	sql.WriteString("SELECT ")

	if q.SelectColumns.Len() > 0 {
		sql.WriteString(q.SelectColumns.String())
	}

	if q.TableName == "" {
		return fmt.Errorf("%w", errors.New("table name found")), ""
	}

	sql.WriteString(fmt.Sprintf(" FROM %s", q.TableName))

	if q.ConditionSQL.Len() > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(q.ConditionSQL.String())
	}

	if q.SQLLimit.Len() > 0 {
		sql.WriteString(" LIMIT ")
		sql.WriteString(q.SQLLimit.String())
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
