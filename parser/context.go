package parser

// ****************************************************************************
// * Kind
// ****************************************************************************

// Kind represents a type of a Value.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	VOID  Kind = `VOID`
	BOOL  Kind = `BOOL`
	INT   Kind = `INT`
	REAL  Kind = `REAL`
	STR   Kind = `STR`
	LIST  Kind = `LIST`
	FUNC  Kind = `FUNC`
	SPELL Kind = `SPELL`
)

// ****************************************************************************
// * Value
// ****************************************************************************

// Value represents a value within the executing program, either the value of
// a variable or an intermediate within a statement.
type Value struct {
	k Kind
	v interface{}
}

// ****************************************************************************
// * Context
// ****************************************************************************

// Context represents the current executing context. It contains all state
// available to the current scope such as available variables and globals. It
// also contains it's parent context so it doubles up as the context stack
// (linked list).
type Context struct {
	vars    map[string]Value
	globals map[string]Value
	parent  *Context
}

// get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx Context) get(id string) (_ Value) {

	v, ok := ctx.vars[id]
	if ok {
		return v
	}

	v, _ = ctx.globals[id]
	return v
}

// set creates or updates a local variable.
func (ctx Context) set(id string, v Value) {
	ctx.vars[id] = v
}
