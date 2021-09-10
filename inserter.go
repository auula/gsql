package gsql

import "reflect"

type Inserter interface {
}

type Into interface {
	Values(v ...interface{}) Builder
}

type Execute struct {
	TableName string
	Columns   []string
	Value     []string
	Obj       interface{}
}

func Insert(model interface{}, filter []string) Into {

	e := &Execute{
		TableName: "",
		Columns:   make([]string, 0, 20),
		Value:     make([]string, 0, 20),
		Obj:       model,
	}

	typeOf := reflect.TypeOf(model)
	e.TableName = typeOf.Name()

	for i := 0; i < typeOf.NumField(); i++ {

		if len(filter) > 0 && filter != nil {
			for _, c := range filter {
				if c == typeOf.Field(i).Tag.Get("db") {
					continue
				}
			}
		}

		e.Columns = append(e.Columns, typeOf.Field(i).Tag.Get("db"))

	}

	return e
}

func (e *Execute) Values(v ...interface{}) Builder {
	panic("implement me")
}
