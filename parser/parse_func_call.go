package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseFuncCall parses a function call.
func (p *Parser) parseFuncCall(id token.Token) Expr {

	c := funcCallExpr{
		id: id,
	}

	p.takeEnsure(token.OPEN_PAREN)
	c.params = p.parseDelimExpr(true)
	p.takeEnsure(token.CLOSE_PAREN)

	return c
}

// funcCallExpr represents an expression for a function call.
type funcCallExpr struct {
	id     token.Token
	params []Expr
}

// String satisfies the Expr interface.
func (ex funcCallExpr) String() (s string) {

	s += "Call (" + ex.id.String() + ")\n"
	s += "\tParams "

	if len(ex.params) == 0 {
		s += "\tN/A"
	} else {

		for _, p := range ex.params {
			s += "\n\t\t" + p.String()
		}

		s += "\n"
	}

	return
}

// Eval satisfies the Expr interface.
func (call funcCallExpr) Eval(ctx Context) Value {

	v := ctx.resolve(call.id.Value)
	if v.k != FUNC {
		panic(bard.NewHorror(call.id, nil,
			"Not a function, variable cannot be invoked",
		))
	}

	f := v.v.(funcValue)

	if argCount := len(call.params) - len(f.input); argCount < 0 {
		panic(bard.NewHorror(call.id, nil, "Not enough arguments"))
	} else if argCount > 0 {
		panic(bard.NewHorror(call.id, nil, "Too many arguments"))
	}

	inSize := len(call.params)
	sub := ctx.sub()

	for _, id := range f.output {
		sub.override(id.Value, Value{VOID, nil}, false)
	}

	for i := 0; i < inSize; i++ {
		val := call.params[i].Eval(ctx)
		id := f.input[i].Value
		sub.override(id, val, false)
	}

	f.body.Eval(sub)

	outSize := len(f.output)
	tuple := make([]Value, outSize)

	for i := 0; i < outSize; i++ {
		val := sub.get(f.output[i].Value)

		if val == (Value{}) {
			val = Value{VOID, nil}
		}

		tuple[i] = val
	}

	return Value{TUPLE, tuple}
}
