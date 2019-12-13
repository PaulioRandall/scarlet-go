package eval

import (
	CTX "github.com/PaulioRandall/scarlet-go/parser/context"
	"github.com/PaulioRandall/scarlet-go/parser/err"
	"github.com/PaulioRandall/scarlet-go/token"
)

// evalSpell creates an Eval function that invokes a spell function when
// called.
func evalSpell(t []token.Token, params []Eval) Eval {
	return func(parent CTX.Context, _ []CTX.Value) (CTX.Value, err.EvalErr) {

		//ctx := parent.Schism()
		return CTX.Value{}, nil
	}
}
