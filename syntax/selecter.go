package syntax

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

type Alias interface {
	As(field interface{}, asName string) string
}

func AS(field interface{}, asName string) string {
	return ""
}
