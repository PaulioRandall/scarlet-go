package eval

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/token"
)

// Expr represents an expression that produces a value when evaluated.
type Expr func(c ctx.Context) (ctx.Value, EvalErr)

// NewForValue creates an Expr that just returns the input value.
func NewForValue(v ctx.Value) Expr {
	return func(ctx.Context) (ctx.Value, EvalErr) {
		return v, nil
	}
}

// NewForID creates an Expr that returns the value associated with a specific ID
// from the context.
func NewForID(t token.Token) Expr {
	return func(c ctx.Context) (ctx.Value, EvalErr) {
		return c.Get(t.Value), nil
	}
}

// NewForListAccess creates an Expr that returns a specified value within a
// list.
func NewForListAccess(idEv, indexEv Expr) Expr {
	return func(c ctx.Context) (ctx.Value, EvalErr) {
		// TODO:
		return ctx.Value{}, nil
	}
}

func NewForOperator(t token.Token) Expr {
	return func(c ctx.Context) (ctx.Value, EvalErr) {
		// TODO:
		return ctx.Value{}, nil
	}
}
