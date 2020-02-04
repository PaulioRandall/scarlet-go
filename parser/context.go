package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
)

// variable represents a value stored against an identifier within a context.
type variable struct {
	val      Value
	isSticky bool
}

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables. It also contains
// it's parent context so it doubles up as the context stack (linked list).
type Context struct {
	vars   map[string]variable
	parent *Context // TODO: Might be obsolete?
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

	appendVars := func(stickies bool) {

		empty := true

		for id, v := range ctx.vars {
			if v.isSticky == stickies {
				empty = false
				s += "\t" + id + " " + string(v.val.k) + ": " + v.val.String() + "\n"
			}
		}

		if empty {
			s += "\t" + "(Empty)" + "\n"
		}

	}

	s += "stickies:" + "\n"
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
func (ctx Context) set(id string, val Value, isSticky bool) {

	if v := ctx.vars[id]; v.isSticky {
		panic(bard.NewNightmare(nil,
			"Cannot reassign the sticky variable '%v'", id,
		))
	}

	ctx.vars[id] = variable{
		val:      val,
		isSticky: isSticky,
	}
}
