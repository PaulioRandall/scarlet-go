package parser2

import (
	"errors"
)

// Kind represents a Value type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	LIST  Kind = `LIST`
	STR   Kind = `STR`
	FUNC  Kind = `FUNC`
	SPELL Kind = `SPELL`
)

// Value represents a value within the script. This could be a variable value
// or intermediate one.
type Value struct {
	k Kind
	v interface{}
}

// ToList returns the value as a list of Values.
func (v Value) ToList() ([]Value, error) {
	if v.k != LIST {
		return nil, errors.New("TODO")
	}
	return v.v.([]Value), nil
}

// ToStr returns the value as a string or error if the kind does not represent
// a string.
func (v Value) ToStr() (string, error) {
	if v.k != STR {
		return ``, errors.New("TODO")
	}
	return v.v.(string), nil
}

// ToFunc returns the value as a function or error if the kind does not
// represent a function.
func (v Value) ToFunc() (Eval, error) {
	if v.k != FUNC {
		return nil, errors.New("TODO")
	}
	return v.v.(Eval), nil
}

// ToSpell returns the value as a spell or error if the kind does not represent
// a spell.
func (v Value) ToSpell() (Eval, error) {
	if v.k != SPELL {
		return nil, errors.New("TODO")
	}
	return v.v.(Eval), nil
}
