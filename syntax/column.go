package syntax

import "fmt"

//// Col("name").Eq(9) name = 9
//// Col(Col("name").Equal(9)).AND(Col("age").Eq(10))

type Columned interface {
	OR(column Columned) Columned
	AND(column Columned) Columned
	Value() string
}

type Compare interface {
	Equal(value interface{}) *Column
}

func (c *Column) Equal(value interface{}) *Column {
	c.value = fmt.Sprintf("%s = %v", c.value, value)
	return c
}

type Column struct {
	op    string
	value string
}

func (c *Column) OR(column Columned) Columned {
	c.value = fmt.Sprintf("%s OR %s", c.value, column.Value())
	return c
}

func (c *Column) AND(column Columned) Columned {
	c.value = fmt.Sprintf("%s AND %s", c.value, column.Value())
	return c
}

func (c *Column) Value() string {
	return c.value
}

func Condition(value Columned) Columned {
	return &Column{
		value: value.Value(),
	}
}

func Col(value string) Compare {
	return &Column{
		value: value,
	}
}
