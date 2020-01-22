package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token"
)

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

// NewValue creates a new value from a token.
func NewValue(tk token.Token) Value {

	var k Kind
	var v interface{}

	switch tk.Kind {
	case token.STR_LITERAL:
		k, v = STR, tk.Value

	case token.BOOL_LITERAL:
		k, v = BOOL, (tk.Value == "TRUE")

	default:
		panic("An UNDEFINED token may not be converted to a Value")
	}

	return Value{
		k: k,
		v: v,
	}
}

// String returns a human readable string representation of the value.
func (v Value) String() string {
	if v.k == STR {
		return "\"" + v.v.(string) + "\""
	}
	return fmt.Sprintf("%v", v.v)
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

// NewContext creates a new context with variable and global identifier maps
// pre-initialised.
func NewContext() Context {
	return Context{
		vars:    make(map[string]Value),
		globals: make(map[string]Value),
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

	s += varsToString("globals", ctx.globals)
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

	v, _ = ctx.globals[id]
	return v
}

// set creates or updates a local variable.
func (ctx Context) set(id string, v Value) {
	ctx.vars[id] = v
}
