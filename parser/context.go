package parser

import (
	"fmt"
	"strconv"

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
	case token.STR_LITERAL, token.STR_TEMPLATE:
		k, v = STR, tk.Value

	case token.BOOL_LITERAL:
		k, v = BOOL, (tk.Value == "TRUE")

	case token.INT_LITERAL:
		k, v = INT, parseNum(tk)

	case token.REAL_LITERAL:
		k, v = REAL, parseNum(tk)

	default:
		panic(tk.String() + ": An UNDEFINED token may not be converted to a Value")
	}

	return Value{
		k: k,
		v: v,
	}
}

// parseNum parses and integer or real token value.
func parseNum(tk token.Token) (v interface{}) {

	var e error

	switch tk.Kind {
	case token.INT_LITERAL:
		v, e = strconv.ParseInt(tk.Value, 10, 64)
	case token.REAL_LITERAL:
		v, e = strconv.ParseFloat(tk.Value, 64)
	default:
		panic(tk.String() + ": Not a number parsable token")
	}

	if e != nil {
		panic(tk.String() + ": Could not parse number token: " + e.Error())
	}

	return
}

// String returns a human readable string representation of the value.
func (v Value) String() string {
	switch v.k {
	case STR:
		return "\"" + v.v.(string) + "\""
	default:
		return fmt.Sprintf("%v", v.v)
	}
}

// ****************************************************************************
// * Context
// ****************************************************************************

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
		panic("Cannot reassign the sticky variable '" + id + "'")
	}

	ctx.stickies[id] = v
}
