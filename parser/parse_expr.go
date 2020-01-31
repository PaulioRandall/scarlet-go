package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseStat parses the next statement.
func (p *Parser) parseStat() Stat {
	switch tk := p.peek(); tk.Kind {
	case token.STICKY, token.ID:
		return p.parseAssign()
	default:
		panic(tk.String() + ": Token does not start a valid expression or " +
			"parsing has not been implemented for it yet")
	}
}

// parseExpr parses the next expression.
func (p *Parser) parseExpr() Expr {

	if p.peek().Kind == token.TERMINATOR {
		p.take()
	}

	switch tk := p.peek(); tk.Kind {
	case token.STR_LITERAL, token.STR_TEMPLATE:
		// TODO: string templates need compiling
		fallthrough
	case token.BOOL_LITERAL, token.INT_LITERAL, token.REAL_LITERAL:
		p.take()
		return valueExpr{
			tokenExpr: tokenExpr{tk},
			v:         NewValue(tk),
		}
	case token.OPEN_LIST:
		return p.parseList()
	default:
		panic(tk.String() + ": Token does not start a valid expression or " +
			"parsing has not been implemented for it yet")
	}
}

// parseDelimExpr parses a delimitered separated set of expressions.
func (p *Parser) parseDelimExpr() (exs []Expr) {

	for {
		if p.peek().Kind == token.CLOSE_LIST {
			return
		}

		ex := p.parseExpr()
		exs = append(exs, ex)

		if p.peek().Kind == token.DELIM {
			p.take()

			if p.peek().Kind == token.TERMINATOR {
				p.take()
			}

			continue
		}

		return
	}
}

// parseList parses a list literal.
func (p *Parser) parseList() Expr {

	start := p.takeEnsure(token.OPEN_LIST)

	if p.peek().Kind == token.TERMINATOR {
		p.take()
	}

	v := p.parseDelimExpr()

	if p.peek().Kind == token.TERMINATOR {
		p.take()
	}

	end := p.takeEnsure(token.CLOSE_LIST)

	return listExpr{
		start: start,
		end:   end,
		items: v,
	}
}
