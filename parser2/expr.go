package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Expr represents an expression. An expression maybe composed of many
// sub expressions but must never recurse.
type Expr interface {

	// Eval executes the expression returning a value. The type of the value
	// returned will depend on the type of expression being executed.
	Eval(ctx Context) Value

	// Token returns the token that links the expression to the source code.
	Token() token.Token

	// String returns the expression as a human readable string.
	String() string
}

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
	return ex.tk.String()
}

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
