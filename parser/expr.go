package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/value"
)

// Context represents a specific scope within a script.
// E.g.
// - Root of the script file
// - Inside a function body `F`
// - Inside a match block `MATCH`
// - etc
type Context interface {
}

// Expression represents a parsed script expression.
type Expression interface {

	// Eval evaluates the expression returning the resultant value.
	Eval(Context) value.Value
}
