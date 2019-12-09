package parser

// Context represents a specific scope within a script.
// E.g.
// - Root of the script file
// - Inside a function body `F`
// - Inside a match block `MATCH`
// - etc
type Context interface {

	// Get returns the value for a specified ID.
	Get(string) Value

	// Set sets the value for a specified ID overwritting any current value.
	Set(string, Value)
}

// RootContext represents a context at the root of a script.
type RootContext struct {
	vars map[string]Value
}

// NewRootCtx creates a new RootContext.
func NewRootCtx() Context {
	return RootContext{
		vars: make(map[string]Value),
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
