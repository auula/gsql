package syntax

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type SortType string

const (
	DESC SortType = "DESC"
	ASC           = "ASC"
)

// Selector statement trait
// Example:
// sql.Selector(user.id,user.name).Where(user.id," > ?",10).From("").Build().
// From("") is empty use model structured name.
type Selector interface {
	Builder
	//Distinct() Selector
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
	OrderBy(row []OrderRow) Filter
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

func Limit(sql Selector, offset bool, index, row int) Builder {

	if strings.Contains(sql.Buf().String(), "LIMIT") {
		sql.Error(errors.New("limit syntax recurring"))
		return sql
	}

	if offset {
		sql.Buf().WriteString(Offset(index, row))
		return sql
	}

	sql.Buf().WriteString(fmt.Sprintf(" LIMIT %v,%v", index, row))
	return sql
}

type OrderRow struct {
	Field string
	Sort  SortType
}

func OrderBy(sql Selector, row []OrderRow) Builder {
	sql.Buf().WriteString(" ORDER BY ")
	for i, iterm := range row {
		sql.Buf().WriteString(fmt.Sprintf(" %s %v", iterm.Field, iterm.Sort))
		if len(row)-1 == i {
			break
		}
		sql.Buf().WriteString(",")
	}
	return sql
}
