package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseStat parses the next statement.
func (p *Parser) parseStat() Stat {
	switch tk := p.peek(); tk.Kind {
	case token.ID:
		return p.parseAssign()
	default:
		panic(tk.String() + ": Token of kind does not start a valid " +
			"statement or parsing has not been implemented for it yet")
	}
}

// parseExpr parses the next expression.
func (p *Parser) parseExpr() Expr {
	switch tk := p.take(); tk.Kind {
	case token.STR_LITERAL, token.STR_TEMPLATE:
		// TODO: string templates need compiling
		fallthrough
	case token.BOOL_LITERAL:
		fallthrough
	case token.INT_LITERAL, token.REAL_LITERAL:
		return valueExpr{
			tokenExpr: tokenExpr{tk},
			v:         NewValue(tk),
		}
	default:
		panic(tk.String() + ": Token of kind does not start a valid " +
			"expression or parsing has not been implemented for it yet")
	}
}
