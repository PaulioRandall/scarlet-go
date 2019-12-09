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

// IdExpr is an expression that resolves an ID into a value
type IdExpr string

// Eval satisfies the Expr interface.
func (ex IdExpr) Eval(ctx Context) Value {
	id := string(ex)
	return ctx.Get(id)
}

// FuncExpr is an expression that calls a function for a value.
type FuncExpr Func

// Eval satisfies the Expr interface.
func (ex FuncExpr) Eval(ctx Context) Value {
	// TODO
	return Value{Func{}}
}

// SpellExpr is an expression that calls a spell for a value.
type SpellExpr Func

// Eval satisfies the Expr interface.
func (ex SpellExpr) Eval(ctx Context) Value {
	// TODO
	return Value{Spell{}}
}
