package value

// Value represents a value within the script. This could be a variable value
// or intermediate one.
type Value struct {
	k Kind
	v interface{}
}
