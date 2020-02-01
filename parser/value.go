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
		k, v = INT, parseNum(INT, tk)

	case token.REAL_LITERAL:
		k, v = REAL, parseNum(REAL, tk)

	default:
		panic(tk.String() + ": An UNDEFINED token may not be converted to a Value")
	}

	return Value{
		k: k,
		v: v,
	}
}

// parseNum parses an INT or REAL token value into its Go counterpart.
func parseNum(k Kind, tk token.Token) (v interface{}) {

	var e error

	if k == INT {
		v, e = strconv.ParseInt(tk.Value, 10, 64)
	} else if k == REAL {
		v, e = strconv.ParseFloat(tk.Value, 64)
	} else {
		// Sanity Check
		panic(tk.String() + ": Illegal number type, cannot parse")
	}

	if e != nil {
		// Sanity Check
		panic(tk.String() + ": Could not parse integer token: " + e.Error())
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
