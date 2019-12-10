package parser2

// Context represents a specific scope within a script. There will be one for
// the root of a script with a new one being created for the body of each
// function when invoked.
type Context interface {

	// Get returns the value for a specified ID. Local variables take precedence
	// over global ones so the local mapping is searched first. An empty Value
	// signifies there was no mapping.
	Get(string) Value

	// Set sets the value for a specified variable overwritting any current value.
	Set(string, Value)

	// SetGlobal sets the value for a specified global variable overwritting any
	// current value.
	SetGlobal(string, Value)

	// Schism creates a new context containing only the global ones of its parent.
	Schism() Context
}

// RootContext represents a context at the root of a script.
type RootContext struct {
	vars    map[string]Value
	globals map[string]Value
}

// NewRootCtx creates a new RootContext.
func NewRootCtx() Context {
	return RootContext{
		vars:    make(map[string]Value),
		globals: make(map[string]Value),
	}
}

// Get satisfies the Context interface.
func (ctx RootContext) Get(id string) Value {
	return ctx.vars[id]
}

// Set satisfies the Context interface.
func (ctx RootContext) Set(id string, v Value) {
	ctx.vars[id] = v
}

// SetGlobal satisfies the Context interface.
func (ctx RootContext) SetGlobal(id string, v Value) {
	ctx.globals[id] = v
}

// Schism satisfies the Context interface.
func (ctx RootContext) Schism() Context {
	return RootContext{
		vars:    make(map[string]Value),
		globals: ctx.globals,
	}
}
