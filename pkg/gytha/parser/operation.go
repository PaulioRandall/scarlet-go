package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func parseOperation(p *pipe, left Expression, leftPriority int) Expression {
	// pattern := {operator operand}
	// operand := literal | expression

	// Warning: this is where parsing gets a little complicated!!

	if p.peek() == nil {
		return left
	}

	op := p.peek()
	m := op.Type()

	if leftPriority >= m.Precedence() {
		// Any token that is not an operator has a precedence of zero, so the left
		// hand expression will always be returned in such a case.
		return left
	}

	// Because operator not taken yet.
	p.expect(`parseOperation`, m)

	// Parse the terminal or expression on the right.
	right := parseSubOperation(p)

	// Recursively parse the expression on the right until an operator with
	// precedence less or equal to this one is encountered.
	right = parseOperation(p, right, m.Precedence())

	left = Operation{left, op, right}

	// Parse the remaining operations in this expression.
	left = parseOperation(p, left, leftPriority)
	return left
}

func parseSubOperation(p *pipe) Expression {
	// pattern := func_call | literal | group

	switch {
	case isNegation(p):
		return parseNegation(p)

	case isSpellCall(p):
		return parseSpellCall(p)

	case isFuncCall(p):
		return parseFuncCall(p)

	case isLiteral(p):
		return parseLiteral(p)

	case isGroup(p):
		return parseGroup(p)
	}

	err.Panic(
		errMsg("parseRightSide", `function call, literal, or group`, p.peek()),
		err.At(p.peek()),
	)

	return nil
}

func isLiteral(p *pipe) bool {
	return p.matchAny(
		TK_IDENTIFIER, // Yes TK_I know, need to sort it out
		TK_VOID,
		TK_BOOL,
		TK_NUMBER,
		TK_STRING,
	)
}

func parseLiteral(p *pipe) Expression {
	tk := p.next()

	if tk.Type() == TK_IDENTIFIER {
		return Identifier{tk}
	} else {
		return Value{tk}
	}
}

func isGroup(p *pipe) bool {
	return p.match(TK_PAREN_OPEN)
}

func parseGroup(p *pipe) Expression {
	// pattern := PAREN_OPEN expression PAREN_CLOSE

	p.expect(`parseGroup`, TK_PAREN_OPEN)

	g := parseExpression(p)
	if g == nil {
		err.Panic(
			errMsg("parseGroup", TK_ANOTHER.String(), p.past()),
			err.At(p.past()),
		)
	}

	p.expect(`parseGroup`, TK_PAREN_CLOSE)
	return g
}

func isNegation(p *pipe) bool {
	return p.match(TK_MINUS)
}

func parseNegation(p *pipe) Expression {

	// pattern := func_call | list_access | literal | group

	n := Negation{
		Tk: p.expect(`parseNegation`, TK_MINUS),
	}

	switch {
	case isFuncCall(p):
		n.Expr = parseFuncCall(p)

	case isListAccess(p):
		n.Expr = parseListAccess(p)

	case isLiteral(p):
		n.Expr = parseLiteral(p)

	case isGroup(p):
		n.Expr = parseGroup(p)

	default:
		err.Panic(errMsg("parseNegation", `term`, p.peek()), err.At(p.peek()))
	}

	return n
}
