package parser

import (
	"errors"
)

// Kind represents a type of value.
type Kind string

const (
	STR Kind = `STRING`
)

// Value represents a value within the script. This could be a variable value
// or intermediate one.
type Value struct {
	k Kind
	v interface{}
}

// ToStr returns the value as a string or error if the kind does not represent
// a string.
func (v Value) ToStr() (string, error) {
	if v.k != STR {
		return ``, errors.New("")
	}
	return v.v.(string), nil
}
