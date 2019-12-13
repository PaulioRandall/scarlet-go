package eval

import (
	CTX "github.com/PaulioRandall/scarlet-go/parser/context"
	"github.com/PaulioRandall/scarlet-go/parser/err"
)

// evalFunc creates an Eval function that invokes a Scarlet function when
// called.
func evalFunc(fn Eval, params []Eval) Eval {
	return func(parent CTX.Context, _ []CTX.Value) (CTX.Value, err.EvalErr) {

		fParams, e := evalParams(parent, params)
		if e != nil {
			return CTX.Value{}, e
		}

		fValue, e := fn(parent, nil)
		if e != nil {
			return CTX.Value{}, e
		}

		f, eerr := fValue.ToFunc()
		if eerr != nil {
			return CTX.Value{}, err.NewEvalErr(eerr, -1, "TODO")
		}

		ctx := parent.Schism()
		v, perr := f(ctx, fParams)
		if perr != nil {
			return CTX.Value{}, err.NewEvalErr(perr, -1, "TODO")
		}

		return v, nil
	}
}
