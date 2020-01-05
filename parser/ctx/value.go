package ctx

import (
	"errors"
)

// Kind represents a type of a Value.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	LIST  Kind = `LIST`
	BOOL  Kind = `BOOL`
	INT   Kind = `INT`
	REAL  Kind = `REAL`
	STR   Kind = `STR`
	FUNC  Kind = `FUNC`
	SPELL Kind = `SPELL`
)

// Procedure is a function prototype that executes a construct with a block of
// code.
type Procedure func(ctx Context, params []Value) (Value, error)

// Value represents a value within the script. This could be a variable value
// or intermediate one.
type Value struct {
	k Kind
	v interface{}
}

// NewValue creates a new Value.
func NewValue(k Kind, v interface{}) Value {
	return Value{k, v}
}

// IsZero returns true if the value is empty.
func (v Value) IsZero() bool {
	return v == (Value{})
}

// ToList returns the value as a list of Values.
func (v Value) ToList() ([]Value, error) {
	if v.k != LIST {
		return nil, errors.New("TODO")
	}
	return v.v.([]Value), nil
}

// ToBool returns the value as a bool or error if the kind does not represent
// a boolean.
func (v Value) ToBool() (bool, error) {
	if v.k != BOOL {
		return false, errors.New("TODO")
	}
	return v.v.(bool), nil
}

// ToInt returns the value as an integer or error if the kind does not represent
// an integer.
func (v Value) ToInt() (int64, error) {
	if v.k != INT {
		return 0, errors.New("TODO")
	}
	return v.v.(int64), nil
}

// ToReal returns the value as an real number or error if the kind does not
// represent a real number.
func (v Value) ToReal() (float64, error) {
	if v.k != REAL {
		return 0, errors.New("TODO")
	}
	return v.v.(float64), nil
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
