package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// funcValueExpr represents a function as a Value.
type funcValueExpr struct {
	input  []token.Token
	output []token.Token
	body   Stat
}

// String
func (ex funcValueExpr) String() (s string) {

	s += "F("

	for i, id := range ex.input {
		if i != 0 {
			s += ", "
		}

		s += id.Value
	}

	if len(ex.output) > 0 {
		s += " -> "

		for i, id := range ex.output {
			if i != 0 {
				s += ", "
			}

			s += id.Value
		}
	}

	return s + ")"
}

/*
// Eval satisfies the Expr interface.
func (ex funcValueExpr) Eval(ctx Context) (_ Value) {

	var val Value
	sub := ctx.sub()

	for _, param := range ex.input {
		val = param.ex.Eval(ctx)
		sub.set(param.id.Value, val, false)
	}

	ex.body.Eval(sub)

	for _, out := range ex.output {
		val = sub.get(out.Value)

		if val == (Value{}) {
			ctx.set(out.Value, Value{VOID, nil}, false)
			continue
		}

		ctx.set(out.Value, val, false)
	}

	return
}

*/
