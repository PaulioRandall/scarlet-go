package parser

import (
	"errors"
)

// Kind represents a type of value.
type Kind string

const (
	STR   Kind = `STRING`
	FUNC  Kind = `FUNC`
	SPELL Kind = `SPELL`
)

// Value represents a value within the script. This could be a variable value
// or intermediate one.
type Value struct {
	k Kind
	v interface{}
}

// NewValue creates a new value.
func NewValue(k Kind, v interface{}) Value {
	return Value{
		k: k,
		v: v,
	}
}

// ToStr returns the value as a string or error if the kind does not represent
// a string.
func (v Value) ToStr() (string, error) {
	if v.k != STR {
		return ``, errors.New("")
	}
	return v.v.(string), nil
}

// ToFunc returns the value as a function or error if the kind does not
// represent a function.
func (v Value) ToFunc() (Func, error) {
	if v.k != FUNC {
		return Func{}, errors.New("")
	}
	return v.v.(Func), nil
}

// ToSpell returns the value as a spell or error if the kind does not represent
// a spell.
func (v Value) ToSpell() (Spell, error) {
	if v.k != SPELL {
		return Spell{}, errors.New("")
	}
	return v.v.(Spell), nil
}
