package syntax

import "fmt"

//// Col("name").Eq(9) name = 9
//// Col(Col("name").Equal(9)).AND(Col("age").Eq(10))

type Columned interface {
	OR(column Compare) Columned
	AND(column Compare) Columned
	String() string
}

type Compare interface {
	String() string
	Equal(value interface{}) Compare
}

func (c *Column) Equal(value interface{}) Compare {
	c.value = fmt.Sprintf("%s = %v", c.value, value)
	return c
}

type Column struct {
	op    string
	value string
}

func (c *Column) OR(column Compare) Columned {
	c.value = fmt.Sprintf("%s OR %s", c.value, column.String())
	return c
}

func (c *Column) AND(column Compare) Columned {
	c.value = fmt.Sprintf("%s AND %s", c.value, column.String())
	return c
}

func (c *Column) String() string {
	return c.value
}

func Condition(value Compare) Columned {
	return &Column{
		value: value.String(),
	}
}

func Col(value string) Compare {
	return &Column{
		value: value,
	}
}
