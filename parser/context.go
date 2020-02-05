package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
)

// variable represents a value stored against an identifier within a context.
type variable struct {
	val     Value
	isFixed bool
}

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables.
type Context struct {
	vars map[string]variable
}

// NewContext creates a new context with variable and global identifier maps
// pre-initialised.
func NewContext() Context {
	return Context{
		vars: make(map[string]variable),
	}
}

// String returns a human readable string representation of the context.
func (ctx Context) String() (s string) {

	appendVars := func(fixed bool) {

		empty := true

		for id, v := range ctx.vars {
			if v.isFixed == fixed {
				empty = false
				s += "\t" + id + " " + string(v.val.k) + ": " + v.val.String() + "\n"
			}
		}

		if empty {
			s += "\t" + "(Empty)" + "\n"
		}

	}

	s += "fixed:" + "\n"
	appendVars(true)

	s += "variables:" + "\n"
	appendVars(false)

	return
}

// get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx Context) get(id string) (_ Value) {

	if v, ok := ctx.vars[id]; ok {
		return v.val
	}

	return
}

// resolve returns the value assigned to a specified variable. If the ID does
// not exist a panic ensues.
func (ctx Context) resolve(id string) (_ Value) {
	v := ctx.get(id)

	if v == (Value{}) {
		panic(bard.NewNightmare(nil,
			"Cannot resolve the variable '%v'", id,
		))
	}

	return v
}

// set creates or updates a variable.
func (ctx Context) set(id string, val Value, isFixed bool) {

	if v := ctx.vars[id]; v.isFixed {
		panic(bard.NewNightmare(nil,
			"Cannot reassign the fixed variable '%v'", id,
		))
	}

	if val.k == VOID {
		delete(ctx.vars, id)
		return
	}

	ctx.vars[id] = variable{
		val:     val,
		isFixed: isFixed,
	}
}

// sub creates a copy of the context without non-sticky variables.
func (ctx Context) sub() Context {

	sub := NewContext()

	for id, val := range ctx.vars {
		if val.isFixed {
			sub.vars[id] = val
		}
	}

	return sub
}
