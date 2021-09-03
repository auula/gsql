package verb

// Selecter select statement trait
type Selecter interface {
	Select() Selecter
	Where() Selecter
	From() Selecter
}
