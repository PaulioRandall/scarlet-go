package parser

// valueExpr represents an expression that simple returns a value. It may also
// have a token which indicates it position within the source code or the
// position of the expression thgat generated it.
type valueExpr struct {
	tokenExpr
	v Value
}

// Eval satisfies the Expr interface.
func (ex valueExpr) Eval(ctx Context) (_ Value) {
	return ex.v
}

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
