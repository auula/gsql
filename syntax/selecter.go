package syntax

import (
	"fmt"
	"reflect"
)

type SortType int8

const (
	DESC SortType = iota
	ASC
)

// Select statement trait
// Example:
// sql.Select(user.id,user.name).Where(user.id," > ?",10).From("").Build().
// From("") is empty use model structured name.
type Select interface {
	Builder
	//Distinct() Select
}

type Form interface {
	From(tab string) Filter
}

type Where interface {
	Where(sql string, v ...interface{}) Filter
}

type Filter interface {
	Where
	Builder
	Limit(offset bool, index int, row int) Filter
	Order(field interface{}, sort SortType) Filter
}

type Alias_ interface {
	As(field interface{}, asName string) string
}

func As(field string, asName string) string {
	return fmt.Sprintf("%s AS '%s'", field, asName)
}

func Alias(model interface{}, aliasMap map[string]string) []string {
	values := make([]string, 0)
	ty := reflect.TypeOf(model)
	for i := 0; i < ty.NumField(); i++ {
		if v, ok := aliasMap[ty.Field(i).Tag.Get("json")]; ok {
			values = append(values, fmt.Sprintf("%s AS '%s'", ty.Field(i).Tag.Get("json"), v))
		} else {
			values = append(values, ty.Field(i).Tag.Get("json"))
		}

	}
	return values
}

func Offset(index, row int) string {
	return fmt.Sprintf(" LIMIT %v OFFSET %v", row, index)
}

func Limit(sql Select, offset bool, index, row int) Builder {
	if offset {
		sql.Buf().WriteString(Offset(index, row))
	}
	sql.Buf().WriteString(fmt.Sprintf(" LIMIT %v,%v", index, row))
	return sql
}
