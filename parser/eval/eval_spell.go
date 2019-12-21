package eval

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/token2"
)

// evalSpell creates an Expr that invokes a spell function when called.
func evalSpell(t []token.Token, params []Expr) Expr {
	return func(parent ctx.Context) (ctx.Value, EvalErr) {

		//ctx := parent.Schism()
		return ctx.Value{}, nil
	}
}
