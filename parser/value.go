package parser

import (
	"errors"
)

// Value represents a value within the script. This could be a variable value
// or intermediate one.
type Value struct {
	v interface{}
}

// ToStr returns the value as a string or error if the kind does not represent
// a string.
func (val Value) ToStr() (string, error) {
	if v, ok := val.v.(string); ok {
		return v, nil
	}
	return ``, errors.New("")
}

// ToFunc returns the value as a function or error if the kind does not
// represent a function.
func (val Value) ToFunc() (Func, error) {
	if v, ok := val.v.(Func); ok {
		return v, nil
	}
	return Func{}, errors.New("")
}

// ToSpell returns the value as a spell or error if the kind does not represent
// a spell.
func (val Value) ToSpell() (Spell, error) {
	if v, ok := val.v.(Spell); ok {
		return v, nil
	}
	return Spell{}, errors.New("")
}
