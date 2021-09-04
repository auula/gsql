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
	Where
	//Distinct() Select
}

type Form interface {
	From(tab string) Select
}

type Where interface {
	Where(sql string, v ...interface{}) Filter
}

type Filter interface {
	Builder
	Limit(x int, y int) Filter
	Order(field interface{}, sort SortType) Filter
}

type Alias_ interface {
	As(field interface{}, asName string) string
}

func As(field string, asName string) string {
	return fmt.Sprintf("%s AS '%s'", field, asName)
}

func Alias(model interface{}, asmap map[string]string) []string {
	values := make([]string, 0)
	ty := reflect.TypeOf(model)
	for i := 0; i < ty.NumField(); i++ {
		if v, ok := asmap[ty.Field(i).Tag.Get("json")]; ok {
			values = append(values, fmt.Sprintf("%s AS '%s'", ty.Field(i).Tag.Get("json"), v))
		} else {
			values = append(values, ty.Field(i).Tag.Get("json"))
		}

	}
	return values
}
