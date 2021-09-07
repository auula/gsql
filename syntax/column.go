package syntax

import (
	"fmt"
	"strings"
)

//// Col("name").Eq(9) name = 9
//// Col(Col("name").Equal(9)).AND(Col("age").Eq(10))

type Columned interface {
	OR(column Compare) Columned
	AND(column Compare) Columned
	String() string
}

// 这里可以指定一下类型

type Compare interface {
	String() string
	In(value []string) Compare
	Like(value string) Compare
	Equal(value interface{}) Compare
	Time(value string) Compare
	Between(value []string) Compare
}

type Column struct {
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

// 这里也是可以的

func Col(value string) Compare {
	return &Column{
		value: value,
	}
}

func (c *Column) In(value []string) Compare {
	buf := new(strings.Builder)
	for i, s := range value {
		buf.WriteString(fmt.Sprintf("'%v'", s))
		if len(value)-1 == i {
			break
		}
		buf.WriteString(",")
	}
	c.value = fmt.Sprintf("%s IN (%v)", c.value, buf)
	return c
}

func (c *Column) Like(value string) Compare {
	c.value = fmt.Sprintf("%s LIKE '%v'", c.value, value)
	return c
}

func (c *Column) Time(value string) Compare {
	c.value = fmt.Sprintf("%s = %v", c.value, value)
	return c
}

func (c *Column) Between(value []string) Compare {
	c.value = fmt.Sprintf("%s = %v", c.value, value)
	return c
}

func (c *Column) Equal(value interface{}) Compare {
	c.value = fmt.Sprintf("%s = %v", c.value, value)
	return c
}
