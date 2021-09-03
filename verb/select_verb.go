package verb

type SortType int8

const (
	DESC SortType = iota
	ASC
)

// Selecter select statement trait
// Exmaple:
// gsql.Select(user.id,user.name).Where(user.id," > ?",10).From("").Build().
// From("") is empty use model structured name.
type Selecter interface {
	Builder
	Filter(filter Filter) Selecter
	Distinct() Selecter
	Select() Selecter
	Where() Selecter
	From() Selecter
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
