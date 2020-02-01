package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

// ****************************************************************************
// * Expr
// ****************************************************************************

// Expr represents an expression. An expression maybe composed of many
// sub expressions but must never recurse.
type Expr interface {

	// Token returns the token that links the expression to the source code.
	Token() token.Token

	// String returns the expression as a human readable string.
	String() string

	// TabString returns the expression as a human readable string but allows for
	// any number of tabs to be placed before each line.
	TabString(int) string

	// Eval executes the expression returning a value. The type of the value
	// returned will depend on the type of expression being executed.
	Eval(ctx Context) Value
}

// Stat (Statement) is an expression that always returns an empty value. The
// difference is purely semantic, i.e. tells the reader of this code that the
// value should always be ignored.
type Stat Expr

// ****************************************************************************
// * tokenExpr
// ****************************************************************************

// tokenExpr is a base structure for expressions that may have an associated
// token linking them to the source code.
type tokenExpr struct {
	tk token.Token
}

// Token satisfies the Expr interface.
func (ex tokenExpr) Token() token.Token {
	return ex.tk
}

// String satisfies the Expr interface.
func (ex tokenExpr) String() string {
	return ex.TabString(0)
}

// TabString satisfies the Expr interface.
func (ex tokenExpr) TabString(tabs int) (s string) {
	return strings.Repeat("\t", tabs) + ex.tk.String()
}

// ****************************************************************************
// * derivedExpr
// ****************************************************************************

// derivedExpr is a base structure for expressions that do not have an
// associated token linking them to the source code.
type derivedExpr struct {
}

// Token satisfies the Expr interface.
func (ex derivedExpr) Token() (_ token.Token) {
	return
}

// String satisfies the Expr interface.
func (ex derivedExpr) String() (_ string) {
	return
}

// ****************************************************************************
// * valueExpr
// ****************************************************************************

// valueExpr represents an expression that simple returns a value.
type valueExpr struct {
	tokenExpr
	v Value
}

// Eval satisfies the Expr interface.
func (ex valueExpr) Eval(ctx Context) (_ Value) {
	return ex.v
}

// ****************************************************************************
// * idExpr
// ****************************************************************************

// idExpr represents an expression that simple returns the value assigned to a
// variable.
type idExpr struct {
	tokenExpr
	id string
}

// Eval satisfies the Expr interface.
func (ex idExpr) Eval(ctx Context) (_ Value) {
	return ctx.get(ex.id)
}
