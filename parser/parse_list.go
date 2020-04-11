package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// parseList parses a list literal.
func (p *Parser) parseList() Expr {

	start := p.takeEnsure(lexeme.LEXEME_LIST_OPEN)
	v := p.parseDelimExpr(false)
	end := p.takeEnsure(lexeme.LEXEME_LIST_CLOSE)

	return listExpr{
		start: start,
		end:   end,
		items: v,
	}
}

// listExpr represents an expression that returns a list value.
type listExpr struct {
	start lexeme.Token
	end   lexeme.Token
	items []Expr
}

// String satisfies the Expr interface.
func (ex listExpr) String() (s string) {

	size := len(ex.items)
	s = "List {"

	for i := 0; i < size; i++ {
		s += "\n" + ex.items[i].String()
	}

	s = strings.ReplaceAll(s, "\n", "\n\t")
	return s + "\n}"
}

// Eval satisfies the Expr interface.
func (ex listExpr) Eval(ctx Context) (_ Value) {

	r := []Value{}

	for _, itemExpr := range ex.items {
		v := itemExpr.Eval(ctx)
		r = append(r, v)
	}

	return Value{
		k: LIST,
		v: r,
	}
}
