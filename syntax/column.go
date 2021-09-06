package syntax

//
//
//// Col("name").Eq(9) name = 9
//// Col(Col("name").Eq(9)).AND(Col("age").Eq(10))
//
//type Columned interface {
//	OR(column interface{}) Columned
//	Col(column interface{}) Columned
//	AND(column interface{}) Columned
//}
//
//type Compare interface {
//	Equal(columned interface{}) Column
//}
//
//type Column struct {
//	op     string
//	value  string
//	column interface{}
//}
//
//func (c *Column) OR(column interface{}) Columned {
//	c.op = "="
//	c.value = column
//}
//
//func (c *Column) Col(column interface{}) Columned {
//	c.column = column
//	c.op = ""
//	c.value = column
//}
//
//func (c *Column) AND(column interface{}) Columned {
//	panic("implement me")
//}
//
//func Col() Columned {
//	return &Column{}
//}
