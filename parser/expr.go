package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token"
)

// Expr represents a parsed expression.
type Expr interface {

	// Tokens returns the tokens that make up the expression.
	Tokens() []token.Token

	// Eval evaluates the expression returning the resultant value.
	Eval(Context) (Value, token.Perror)
}

// ValueExpr is an expression that just returns itself as a value.
type ValueExpr struct {
	t []token.Token
	v Value
}

// Tokens satisfies the Expr interface.
func (ex ValueExpr) Tokens() []token.Token {
	return ex.t
}

// Eval satisfies the Expr interface.
func (ex ValueExpr) Eval(ctx Context) (Value, token.Perror) {
	return Value(ex.v), nil
}

// IdExpr is an expression that resolves an ID into a value
type IdExpr struct {
	t  []token.Token
	id string
}

// Tokens satisfies the Expr interface.
func (ex IdExpr) Tokens() []token.Token {
	return ex.t
}

// Eval satisfies the Expr interface.
func (ex IdExpr) Eval(ctx Context) (Value, token.Perror) {
	return ctx.Get(ex.id), nil
}

// FuncExpr is an expression that calls a function for a value.
type FuncExpr struct {
	t []token.Token
	f Func
}

// Tokens satisfies the Expr interface.
func (ex FuncExpr) Tokens() []token.Token {
	return ex.t
}

// Eval satisfies the Expr interface.
func (ex FuncExpr) Eval(ctx Context) (Value, token.Perror) {
	// TODO
	return Value{Func{}}, nil
}

// SpellExpr is an expression that calls a spell for a value.
type SpellExpr struct {
	t []token.Token
	p Params
}

// Tokens satisfies the Expr interface.
func (ex SpellExpr) Tokens() []token.Token {
	return ex.t
}

// Eval satisfies the Expr interface.
func (ex SpellExpr) Eval(ctx Context) (Value, token.Perror) {
	// TODO
	return Value{Spell{}}, nil
}

// AssignExpr is an expression that assigns the result of some expressions to
// some variables.
type AssignExpr struct {
	t []token.Token
	a Assign
}

// Tokens satisfies the Expr interface.
func (ex AssignExpr) Tokens() []token.Token {
	return ex.t
}

// Eval satisfies the Expr interface.
func (ex AssignExpr) Eval(ctx Context) (_ Value, e token.Perror) {

	vals := []Value{}
	var v Value

	for _, src := range ex.a.Src {
		v, e = src.Eval(ctx)
		if e != nil {
			return
		}

		vals = append(vals, v)
	}

	if len(vals) != len(ex.a.Dst) {
		e = newPerror(
			ex,
			"There are %d IDs, but the right side expressions resolved to %d values",
			len(ex.a.Dst), len(vals),
		)
		return
	}

	for i, id := range ex.a.Dst {
		ctx.Set(id, vals[i])
	}

	return
}

func newPerror(ex Expr, s string, p ...interface{}) token.Perror {

	t := ex.Tokens()

	return token.NewPerror(
		fmt.Sprint(s, p),
		t[0].Where().Line(),
		t[0].Where().Start(),
		t[len(t)-1].Where().End(),
	)
}
