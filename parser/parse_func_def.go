package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

// parseFuncDef parses a function definition.
func (p *Parser) parseFuncDef() Expr {

	f := funcDefExpr{
		opener: p.takeEnsure(token.KIND_FUNC),
	}

	p.takeEnsure(token.KIND_OPEN_PAREN)

	if p.peek().Kind != token.KIND_CLOSE_PAREN {
		if p.peek().Kind != token.KIND_RETURNS {
			f.input = p.parseIDs()
		}

		if p.peek().Kind == token.KIND_RETURNS {
			p.take()
			f.output = p.parseIDs()
		}
	}

	closeParen := p.takeEnsure(token.KIND_CLOSE_PAREN)

	if p.peek().Kind == token.KIND_DO {
		f.body = p.parseStats(p.take())
	} else {
		f.body = blockStat{
			opener: token.New(token.KIND_INLINE, "", closeParen.Line, closeParen.Col),
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

// String satisfies the Expr interface.
func (ex funcDefExpr) String() (s string) {

	s += "Func (" + ex.opener.String() + ")"

	if len(ex.input) > 0 {
		s += "\n\tIn"

		for _, id := range ex.input {
			s += "\n\t\t" + id.String()
		}
	}

	if len(ex.output) > 0 {
		s += "\n\tOut"

		for _, id := range ex.output {
			s += "\n\t\t" + id.String()
		}
	}

	return s + "\n\t" + strings.ReplaceAll(ex.body.String(), "\n", "\n\t")
}

// Eval satisfies the Expr interface.
func (ex funcDefExpr) Eval(_ Context) (_ Value) {
	return Value{
		k: FUNC,
		v: funcValue{
			input:  ex.input,
			output: ex.output,
			body:   ex.body,
		},
	}
}

// funcValue represents a function as a Value.
type funcValue struct {
	input  []token.Token
	output []token.Token
	body   Stat
}

// String
func (ex funcValue) String() (s string) {

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
