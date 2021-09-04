package syntax

// Builder generate structured query language code string
type Builder interface {
	Build() (error, string)
	String() string
}
