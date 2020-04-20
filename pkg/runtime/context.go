package runtime

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables.
type Context struct {
	vars   map[string]Value
	parent *Context
}

// String returns a human readable string representation of the context.
func (ctx *Context) String() (s string) {

	const NEWLINE = "\n"
	const TAB = "\t"

	s += "variables:" + NEWLINE

	if len(ctx.vars) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return
	}

	for id, v := range ctx.vars {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return
}

// get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx *Context) Get(id string) Value {
	v, ok := ctx.vars[id]

	if !ok {
		v = Void{}
	}

	return v
}

// set creates or updates a variable. Passing a VOID value deletes the entry if
// it exists.
func (ctx *Context) Set(id string, v Value) {

	if _, ok := v.(Void); ok {
		delete(ctx.vars, id)
		return
	}

	ctx.vars[id] = v
}

func (ctx *Context) Spawn() *Context {
	return &Context{
		vars:   make(map[string]Value),
		parent: ctx,
	}
}

func (ctx *Context) Parent() *Context {
	return ctx.parent
}
