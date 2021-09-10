# GSQL

GSQL is a structured query language code builder for golang.

## Genreate SQL

1. Model Structured

```go
type UserInfo struct {
	Id   int    `db:"id" pk:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
```

`db` identifies the database table field，`pk` is the primary key.

2. Simple query

```go
sql1 := gsql.Select().From(UserInfo{})

// SELECT id, name, age FROM UserInfo
t.Log(sql1)
```


3. Alias and pass primary key

```go

sql2 := gsql.SelectAs([]string{"name", gsql.As("age", "年龄"), "id"}).From(UserInfo{}).ById(2)

// SELECT id, name, age AS '年龄', id FROM UserInfo WHERE id = 2
t.Log(sql2)


sql3 := gsql.SelectAs(gsql.Alias(UserInfo{}, map[string]string{
    "name": "名字",
})).From(UserInfo{}).ById(1)

// SELECT id, name AS '名字', age, id FROM UserInfo WHERE id = 1
t.Log(sql3)
```
`SelectAs` if the parameter is `slice`，`id = pk:"id"` The last parameter must be the primary key field of the table.

4. Query a set of data to provide the primary key

```go
sql := gsql.Select().From(UserInfo{}).ByIds(1, 2, 3)

// SELECT id, name, age FROM UserInfo WHERE id IN (1, 2, 3)
t.Log(sql)
```

5. Query by `In`
```go
sql := gsql.Select().From(UserInfo{}).In("age", 21, 19, 28)

// SELECT id, name, age FROM UserInfo WHERE age IN (21, 19, 28)
t.Log(sql)
```

6. query a piece of data

```go
func TestSelectOne(t *testing.T) {
	
	// SELECT id, name, age FROM UserInfo LIMIT 1
	_, sql := gsql.Select().From(UserInfo{}).One()
	t.Log(sql)

	// SELECT id, name, age FROM UserInfo WHERE age > 10 LIMIT 1
	err, sql2 := gsql.Select().From(UserInfo{}).Where("age > ?", 10).One()

	if err == nil {
		t.Log(sql2)
	}
}
```

7. Sort or filter
```go
func TestSelectLimit(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE age > 10 LIMIT 3 OFFSET 1
	sql2 := gsql.Select().From(UserInfo{}).Where("age > ?", 10).Limit(true, 1, 3)

	t.Log(sql2)

}

func TestSelectOrder(t *testing.T) {

	// SELECT id, name, age FROM UserInfo WHERE age > 10 ORDER BY id ASC LIMIT 3 OFFSET 1
	sql2 := gsql.Select().From(UserInfo{}).Where("age > ?", 10).Order([]gsql.Rows{
		{"id", "ASC"},
	}).Limit(true, 1, 3)

	t.Log(sql2)

}
```

8. Insert Data to table

```go
func TestInsert(t *testing.T) {
	// INSERT INTO UserInfo (id, name, age) VALUES (1001, 'Tom', 21)
	sql := gsql.Insert(UserInfo{}, nil).Values(1001, "Tom", 21)
	t.Log(sql)
}

func TestInsertFilter(t *testing.T) {
	// INSERT INTO UserInfo (name, age) VALUES ('Tom', 21)
	err, sql := gsql.Insert(UserInfo{}, []string{"id"}).Values("Tom", 21).Build()
	if err != nil {
		t.Log(err)
	}
	t.Log(sql)
}
```
