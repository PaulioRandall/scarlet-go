package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseStat parses the next statement.
func (p *Parser) parseStat(inline bool) Stat {
	switch tk := p.peek(); tk.Kind {
	case token.FIX, token.ID:
		return p.parseAssign(inline)
	default:
		panic(bard.NewHorror(tk, nil,
			"Unexpected token or maybe parsing has not been implemented for it yet",
		))
	}
}

// parseExpr parses the next expression.
func (p *Parser) parseExpr(allowVoids bool) Expr {

	if p.peek().Kind == token.TERMINATOR {
		p.take()
	}

	switch tk := p.peek(); tk.Kind {
	case token.VOID:
		if !allowVoids {
			panic(bard.NewHorror(tk, nil, "Naughty use of VOID expression"))
		}

		return valueExpr{
			tk: p.take(),
			v:  NewValue(tk),
		}
	case token.ID:
		return idExpr{
			tk: p.take(),
			id: tk.Value,
		}
	case token.STR_LITERAL, token.STR_TEMPLATE:
		// TODO: string templates need compiling
		fallthrough
	case token.BOOL_LITERAL, token.INT_LITERAL, token.REAL_LITERAL:
		return valueExpr{
			tk: p.take(),
			v:  NewValue(tk),
		}
	case token.OPEN_LIST:
		return p.parseList()
	case token.FUNC:
		return p.parseFuncDef()
	default:
		panic(bard.NewHorror(tk, nil,
			"Token does not start a valid expression or "+
				"parsing has not been implemented for it yet",
		))
	}
}

// parseDelimExpr parses a delimitered separated set of expressions.
func (p *Parser) parseDelimExpr(allowVoids bool) (exs []Expr) {

	for p.peek().Kind != token.CLOSE_LIST {

		ex := p.parseExpr(allowVoids)
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

	return
}
