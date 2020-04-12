package parser

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// parseFuncCall parses a function call.
func (p *Parser) parseFuncCall(id lexeme.Token) Expr {

	c := funcCallExpr{
		id: id,
	}

	p.takeEnsure(lexeme.LEXEME_PAREN_OPEN)
	c.params = p.parseDelimExpr(true)
	p.takeEnsure(lexeme.LEXEME_PAREN_CLOSE)

	return c
}

// funcCallExpr represents an expression for a function call.
type funcCallExpr struct {
	id     lexeme.Token
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
	}

	return
}

// Eval satisfies the Expr interface.
func (call funcCallExpr) Eval(ctx Context) Value {

	v := ctx.resolve(call.id.Value)
	if v.k != FUNC {
		panic(newTkErr(call.id,
			"Not a function, variable cannot be invoked",
		))
	}

	f := v.v.(funcValue)

	if argCount := len(call.params) - len(f.input); argCount < 0 {
		panic(newTkErr(call.id, "Not enough arguments"))
	} else if argCount > 0 {
		panic(newTkErr(call.id, "Too many arguments"))
	}

	paramCount := len(call.params)
	sub := NewContext()

	for _, id := range f.output {
		sub.set(id.Value, Value{VOID, nil})
	}

	for i := 0; i < paramCount; i++ {
		val := call.params[i].Eval(ctx)
		id := f.input[i].Value
		sub.set(id, val)
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
