package parser2

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Eval is a function that produces a value when invoked.
type Eval func(Context) (Value, token.Perror)

// evalValue creates an Eval function that just returns the input value.
func evalValue(v Value) Eval {
	return func(Context) (Value, token.Perror) {
		return v, nil
	}
}

// evalID creates an Eval function that returns the value associated with a
// specific ID from the context.
func evalID(t token.Token) Eval {
	return func(ctx Context) (Value, token.Perror) {
		return ctx.Get(t.Value()), nil
	}
}

// evalFunc creates an Eval function that invokes a Scarlet function when
// called.
func evalFunc(t []token.Token, params []Eval) Eval {
	return func(parent Context) (Value, token.Perror) {

		//ctx := parent.Schism()
		return Value{}, nil
	}
}

// evalSpell creates an Eval function that invokes a spell function when
// called.
func evalSpell(t []token.Token, params []Eval) Eval {
	return func(parent Context) (Value, token.Perror) {

		//ctx := parent.Schism()
		return Value{}, nil
	}
}

func evalAssign() Eval {
	return nil
}
