package eval

import (
	CTX "github.com/PaulioRandall/scarlet-go/parser/context"
	"github.com/PaulioRandall/scarlet-go/token"
)

// Eval is a function that produces a value when invoked.
type Eval func(ctx CTX.Context, params []CTX.Value) (CTX.Value, EvalErr)

// evalValue creates an Eval function that just returns the input value.
func evalValue(v CTX.Value) Eval {
	return func(CTX.Context, []CTX.Value) (CTX.Value, EvalErr) {
		return v, nil
	}
}

// evalID creates an Eval function that returns the value associated with a
// specific ID from the context.
func evalID(t token.Token) Eval {
	return func(ctx CTX.Context, _ []CTX.Value) (CTX.Value, EvalErr) {
		return ctx.Get(t.Value()), nil
	}
}

// evalParams executes a set of Eval functions that return values used as
// parameters to some other expression.
func evalParams(ctx CTX.Context, params []Eval) ([]CTX.Value, EvalErr) {

	var r []CTX.Value

	for _, p := range params {
		v, e := p(ctx, nil)

		if e != nil {
			return nil, e
		}

		if v == (CTX.Value{}) {
			return nil, NewEvalErr(nil, -1, "TODO")
		}

		r = append(r, v)
	}

	return r, nil
}
