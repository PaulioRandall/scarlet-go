package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// ****************************************************************************
// * Expr
// ****************************************************************************

// Expr represents an expression. An expression maybe composed of many
// sub expressions but must never recurse.
type Expr interface {

	// String returns the expression as a human readable string.
	String() string

	// Eval executes the expression returning a value. The type of the value
	// returned will depend on the type of expression being executed.
	Eval(Context) Value
}

// Stat (Statement) is an expression that always returns an empty value. The
// difference is purely semantic, i.e. tells the reader of this code that the
// value should always be ignored.
type Stat Expr

// ****************************************************************************
// * valueExpr
// ****************************************************************************

// valueExpr represents an expression that simple returns a value.
type valueExpr struct {
	tk token.Token
	v  Value
}

// Eval satisfies the Expr interface.
func (ex valueExpr) Eval(_ Context) (_ Value) {
	return ex.v
}

// String satisfies the Expr interface.
func (ex valueExpr) String() string {
	return ex.tk.String()
}

// ****************************************************************************
// * idExpr
// ****************************************************************************

// idExpr represents an expression that simple returns the value assigned to a
// variable.
type idExpr struct {
	tk token.Token
	id string
}

// Eval satisfies the Expr interface.
func (ex idExpr) Eval(ctx Context) (_ Value) {
	return ctx.resolve(ex.id)
}

// String satisfies the Expr interface.
func (ex idExpr) String() string {
	return ex.tk.String()
}
