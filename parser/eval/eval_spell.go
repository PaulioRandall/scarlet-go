package eval

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
)

// NewForSpellCall creates an Expr that invokes a spell function when called.
func NewForSpellCall(id Expr, params []Expr) Expr {
	return func(parent ctx.Context) (ctx.Value, EvalErr) {

		//ctx := parent.Schism()
		return ctx.Value{}, nil
	}
}
