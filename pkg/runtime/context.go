package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables.
type Context struct {
	fixed  map[string]Value
	vars   map[string]Value
	parent *Context
}

// String returns a human readable string representation of the context.
func (ctx *Context) String() (s string) {

	const NEWLINE = "\n"
	const TAB = "\t"

	s += "variables:" + NEWLINE

	if len(ctx.vars) == 0 && len(ctx.fixed) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return
	}

	for id, v := range ctx.fixed {
		s += TAB + "FIX " + id + " " + v.String() + NEWLINE
	}

	for id, v := range ctx.vars {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return
}

// Get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx *Context) Get(id string) Value {
	v, ok := ctx.fixed[id]

	if !ok {
		v, ok = ctx.vars[id]

		if !ok {
			v = Void{}
		}
	}

	return v
}

func (ctx *Context) Set(id st.Identifier, v Value) {

	name := id.Source.Value

	if _, ok := ctx.fixed[name]; ok {
		panic(err("Set", id.Token(), "Cannot change a fixed variable"))
	}

	if id.Fixed {
		delete(ctx.vars, name)
		ctx.fixed[name] = v
		return
	}

	if _, ok := v.(Void); ok {
		delete(ctx.vars, name)
		return
	}

	ctx.vars[name] = v
}

func (ctx *Context) Spawn() *Context {

	fixed := make(map[string]Value, len(ctx.fixed))

	for k, v := range ctx.fixed {
		fixed[k] = v
	}

	return &Context{
		fixed:  fixed,
		vars:   make(map[string]Value),
		parent: ctx,
	}
}

func (ctx *Context) Parent() *Context {
	return ctx.parent
}
