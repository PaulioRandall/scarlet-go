package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// Stat (Statement) is an expression that always returns an empty value. The
// difference is purely semantic, i.e. tells the reader of this code that the
// value should always be ignored.
type Stat Expr

// parseStat parses the next statement.
func (p *Parser) parseStat(inline bool) Stat {
	switch tk := p.peek(); tk.Lexeme {
	case lexeme.LEXEME_ID:
		return p.parseAssign(inline)
	default:
		panic(bard.NewHorror(tk, nil,
			"Unexpected token or maybe parsing has not been implemented for it yet",
		))
	}
}

// parseDelimExpr parses a delimitered separated set of expressions.
func (p *Parser) parseDelimExpr(allowVoid bool) (exs []Expr) {

	for p.peek().Lexeme != lexeme.LEXEME_CLOSE_LIST {

		ex := p.parseAssignable(allowVoid)
		exs = append(exs, ex)

		if p.peek().Lexeme == lexeme.LEXEME_DELIM {
			p.take()

			if p.peek().Lexeme == lexeme.LEXEME_TERMINATOR {
				p.take()
			}

			continue
		}

		return
	}

	return
}

// parseExpr parses the an expression that will become an assignment source.
func (p *Parser) parseAssignable(allowVoid bool) Expr {

	if p.peek().Lexeme == lexeme.LEXEME_TERMINATOR {
		p.take()
	}

	switch tk := p.peek(); tk.Lexeme {
	case lexeme.LEXEME_VOID:
		if !allowVoid {
			panic(bard.NewHorror(tk, nil, "Naughty use of void expression"))
		}

		return valueExpr{
			tk: p.take(),
			v:  NewValue(tk),
		}
	case lexeme.LEXEME_FUNC:
		return p.parseFuncDef()
	default:
		return p.parseExpr()
	}
}

// parseExpr parses the next expression.
func (p *Parser) parseExpr() Expr {

	if p.peek().Lexeme == lexeme.LEXEME_TERMINATOR {
		p.take()
	}

	tk := p.peek()
	left := p.parseOperand()

	if left == nil {
		panic(bard.NewHorror(tk, nil,
			"Token does not start a valid expression or "+
				"parsing has not been implemented for it yet",
		))
	}

	if p.identifyOperatorKind(p.peek().Lexeme) == NOT_OPERATOR {
		return left
	}

	return p.parseOperation(left)
}

// parseDelimExpr parses an operand as an expression. Nil is returned if the
// next tokens do not match an operand.
func (p *Parser) parseOperand() (ex Expr) {

	tk := p.peek()

	switch tk.Lexeme {
	case lexeme.LEXEME_STRING, lexeme.LEXEME_TEMPLATE:
		// TODO: string templates need compiling
		fallthrough
	case lexeme.LEXEME_BOOL, lexeme.LEXEME_INT, lexeme.LEXEME_FLOAT:
		ex = valueExpr{
			tk: p.take(),
			v:  NewValue(tk),
		}
	case lexeme.LEXEME_OPEN_LIST:
		ex = p.parseList()
	case lexeme.LEXEME_ID:
		p.take()

		if p.peek().Lexeme == lexeme.LEXEME_OPEN_PAREN {
			ex = p.parseFuncCall(tk)
			break
		}

		ex = idExpr{
			tk: tk,
			id: tk.Value,
		}
	}

	return
}
