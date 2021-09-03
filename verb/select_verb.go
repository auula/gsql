package verb

// Selecter select statement trait
// Exmaple:
// gsql.Select(user.id,user.name).Where(user.id," > ?",10).From("").Build().
// From("") is empty use model structured name.
type Selecter interface {
	Select() Selecter
	Where() Selecter
	From() Selecter
}
