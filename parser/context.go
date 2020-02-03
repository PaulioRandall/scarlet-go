package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
)

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables. It also contains
// it's parent context so it doubles up as the context stack (linked list).
type Context struct {
	stickies map[string]Value
	vars     map[string]Value
	parent   *Context
}

// NewContext creates a new context with variable and global identifier maps
// pre-initialised.
func NewContext() Context {
	return Context{
		stickies: make(map[string]Value),
		vars:     make(map[string]Value),
	}
}

// String returns a human readable string representation of the context.
func (ctx Context) String() (s string) {

	varsToString := func(name string, vars map[string]Value) (s string) {

		s += name + ":" + "\n"

		if len(vars) == 0 {
			s += "\t" + "(Empty)" + "\n"

		} else {
			for k, v := range vars {
				s += "\t" + k + " " + string(v.k) + ": " + v.String() + "\n"
			}
		}

		return s
	}

	s += varsToString("stickies", ctx.stickies)
	s += varsToString("vars", ctx.vars)

	return
}

// get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx Context) get(id string) (_ Value) {

	v, ok := ctx.vars[id]
	if ok {
		return v
	}

	v, _ = ctx.stickies[id]
	return v
}

// set creates or updates a local variable.
func (ctx Context) set(id string, v Value) {
	ctx.vars[id] = v
}

// setSticky creates a sticky variable.
func (ctx Context) setSticky(id string, v Value) {

	if _, exists := ctx.stickies[id]; exists {
		panic(bard.NewNightmare(nil,
			"Cannot reassign the sticky variable '%v'", id,
		))
	}

	ctx.stickies[id] = v
}
