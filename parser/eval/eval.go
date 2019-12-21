package eval

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/token"
)

// Expr represents an expression that produces a value when evaluated.
type Expr func(c ctx.Context) (ctx.Value, EvalErr)

// evalValue creates an Expr that just returns the input value.
func evalValue(v ctx.Value) Expr {
	return func(ctx.Context) (ctx.Value, EvalErr) {
		return v, nil
	}
}

// evalID creates an Expr that returns the value associated with a specific ID
// from the context.
func evalID(t token.Token) Expr {
	return func(c ctx.Context) (ctx.Value, EvalErr) {
		return c.Get(t.Value), nil
	}
}
