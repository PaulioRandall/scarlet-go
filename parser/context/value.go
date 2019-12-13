package context

import (
	"errors"
)

// Kind represents a type of a Value.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	LIST  Kind = `LIST`
	STR   Kind = `STR`
	FUNC  Kind = `FUNC`
	SPELL Kind = `SPELL`
)

// Procedure is a function prototype that executes a construct with a block of
// code.
type Procedure func(ctx Context, params []Value) (Value, ProcErr)

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

// ToFunc returns the value as an executable Scarlet function or error if the
// kind does not represent a function.
func (v Value) ToFunc() (Procedure, error) {
	if v.k != FUNC {
		return nil, errors.New("TODO")
	}
	return v.v.(Procedure), nil
}

// ToSpell returns the value as executable Scarlet spell or error if the kind
// does not represent a spell.
func (v Value) ToSpell() (Procedure, error) {
	if v.k != SPELL {
		return nil, errors.New("TODO")
	}
	return v.v.(Procedure), nil
}
