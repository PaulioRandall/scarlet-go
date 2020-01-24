package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseStat parses the next statement.
func (p *Parser) parseStat(lead token.Token) Stat {
	switch lead.Kind {
	case token.ID:
		return p.parseAssign(lead)
	default:
		panic(lead.String() + ": Token of kind does not start a valid " +
			"statement or parsing has not been implemented for it yet")
	}
}

// parseExpr parses the next expression.
func (p *Parser) parseExpr(lead token.Token) Expr {
	switch lead.Kind {
	case token.STR_LITERAL, token.STR_TEMPLATE:
		// TODO: string templates need compiling
		fallthrough
	case token.BOOL_LITERAL:
		fallthrough
	case token.INT_LITERAL, token.REAL_LITERAL:
		return valueExpr{
			tokenExpr: tokenExpr{lead},
			v:         NewValue(lead),
		}
	default:
		panic(lead.String() + ": Token of kind does not start a valid " +
			"expression or parsing has not been implemented for it yet")
	}
}
