package parser2

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Eval is a function that produces a value when invoked.
type Eval func(ctx Context, params []Value) (Value, EvalErr)

// evalValue creates an Eval function that just returns the input value.
func evalValue(v Value) Eval {
	return func(Context, []Value) (Value, EvalErr) {
		return v, nil
	}
}

// evalID creates an Eval function that returns the value associated with a
// specific ID from the context.
func evalID(t token.Token) Eval {
	return func(ctx Context, _ []Value) (Value, EvalErr) {
		return ctx.Get(t.Value()), nil
	}
}

// evalParams executes a set of Eval functions that return values used as
// parameters to some other expression.
func evalParams(ctx Context, params []Eval) ([]Value, EvalErr) {

	var r []Value

	for _, p := range params {
		v, e := p(ctx, nil)

		if e != nil {
			return nil, e
		}

		if v == (Value{}) {
			return nil, stdEvalErr{
				"TODO",
				-1,
			}
		}

		r = append(r, v)
	}

	return r, nil
}

// evalFunc creates an Eval function that invokes a Scarlet function when
// called.
func evalFunc(fn Eval, params []Eval) Eval {
	return func(parent Context, _ []Value) (Value, EvalErr) {

		fParams, e := evalParams(parent, params)
		if e != nil {
			return Value{}, e
		}

		fValue, e := fn(parent, nil)
		if e != nil {
			return Value{}, e
		}

		f, err := fValue.ToFunc()
		if err != nil {
			return Value{}, stdEvalErr{
				"TODO",
				-1,
			}
		}

		ctx := parent.Schism()
		return f(ctx, fParams)
	}
}

// evalSpell creates an Eval function that invokes a spell function when
// called.
func evalSpell(t []token.Token, params []Eval) Eval {
	return func(parent Context, _ []Value) (Value, EvalErr) {

		//ctx := parent.Schism()
		return Value{}, nil
	}
}

func evalAssign() Eval {
	return nil
}
