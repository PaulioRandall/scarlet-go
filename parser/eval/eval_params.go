package eval

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
)

// evalParams executes a set of expressions that return values used as
// parameters to some other expression.
func evalParams(paramExprs []Expr) Expr {
	return func(c ctx.Context) (r ctx.Value, e EvalErr) {

		var params []ctx.Value
		var v ctx.Value

		for _, expr := range paramExprs {
			v, e = expr(c)

			if e != nil {
				return
			}

			if v.IsEmpty() {
				e = NewEvalErr(nil, -1, "TODO")
				return
			}

			params = append(params, v)
		}

		r = ctx.NewValue(ctx.LIST, params)
		return
	}
}
