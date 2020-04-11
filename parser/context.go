package parser

import (
	"github.com/PaulioRandall/scarlet-go/err"
)

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables.
type Context struct {
	vars map[string]Value
}

// NewContext creates a new context with an intialised variable map.
func NewContext() Context {
	return Context{
		vars: make(map[string]Value),
	}
}

// String returns a human readable string representation of the context.
func (ctx Context) String() (s string) {

	s += "variable:" + "\n"

	if len(ctx.vars) == 0 {
		s += "\t" + "(Empty)" + "\n"
		return
	}

	for id, v := range ctx.vars {
		s += "\t" + id + " (" + string(v.k) + ") " + v.String() + "\n"
	}

	return
}

// get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx Context) get(id string) (_ Value) {

	if v, ok := ctx.vars[id]; ok {
		return v
	}

	return
}

// resolve returns the value assigned to a specified variable. If the ID does
// not exist a panic ensues.
func (ctx Context) resolve(id string) (_ Value) {
	v := ctx.get(id)

	if v == (Value{}) {
		panic(err.NewNightmare(nil,
			"Cannot resolve the variable '%v'", id,
		))
	}

	return v
}

// set creates or updates a variable. Passing a VOID value signifies the entry
// is to be deleted.
func (ctx Context) set(id string, v Value) {

	if v.k == VOID {
		delete(ctx.vars, id)
		return
	}

	ctx.vars[id] = v
}
