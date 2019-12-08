package parser

// Expr represents a parsed expression.
type Expr interface {

	// Eval evaluates the expression returning the resultant value.
	Eval(Context) Value
}

// ValueExpr is an expression that just returns itself as a value.
type ValueExpr Value

// Eval satisfies the Expr interface.
func (ex ValueExpr) Eval(ctx Context) Value {
	return Value(ex)
}
