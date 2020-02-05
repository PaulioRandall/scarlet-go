package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseFuncDef parses a function definition.
func (p *Parser) parseFuncDef() Expr {

	f := funcDefExpr{
		opener: p.takeEnsure(token.FUNC),
	}

	p.takeEnsure(token.OPEN_PAREN)

	if p.peek().Kind != token.CLOSE_PAREN {
		if p.peek().Kind != token.RETURNS {
			f.input = p.parseIDs()
		}

		if p.peek().Kind == token.RETURNS {
			p.take()
			f.output = p.parseIDs()
		}
	}

	closeParen := p.takeEnsure(token.CLOSE_PAREN)

	if p.peek().Kind == token.DO {
		f.body = p.parseStats(p.take())
	} else {
		f.body = blockStat{
			opener: token.New(token.INLINE, "", closeParen.Line, closeParen.Col),
			block:  []Stat{p.parseStat(true)},
		}
	}

	return f
}

// funcDefExpr represents an expression for a function definition, i.e. an
// expression which creates a function.
type funcDefExpr struct {
	opener token.Token
	input  []token.Token
	output []token.Token
	body   Stat
}

// Token satisfies the Expr interface.
func (ex funcDefExpr) Token() token.Token {
	return ex.opener
}

// String satisfies the Expr interface.
func (ex funcDefExpr) String() (s string) {

	s += ex.opener.String() + "\n"

	if len(ex.input) > 0 {
		s += "~In  "

		for _, id := range ex.input {
			s += "[" + id.String() + "] "
		}

		s += "\n"
	}

	if len(ex.output) > 0 {
		s += "~Out "

		for _, id := range ex.output {
			s += "[" + id.String() + "] "
		}

		s += "\n"
	}

	return s + "~" + ex.body.String()
}

// Eval satisfies the Expr interface.
func (ex funcDefExpr) Eval(_ Context) (_ Value) {

	// TODO: Create a specific struct for the value.
	return Value{
		k: FUNC,
		v: ex,
	}
}

/*


function call


// Eval satisfies the Expr interface.
func (ex funcDefExpr) Eval(ctx Context) (_ Value) {

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
