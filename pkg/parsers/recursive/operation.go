package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func parseOperation(p *pipe, left st.Expression, leftPriority int) st.Expression {
	// pattern := {operator operand}
	// operand := literal | expression

	// Warning: this is where parsing gets a little complicated!!

	op := p.peek()

	if leftPriority >= op.Precedence() {
		// Any token that is not an operator has a precedence of zero, so the left
		// hand expression will always be returned in such a case.
		return left
	}

	// Because operator not taken yet.
	p.expect(`parseOperation`, op.Type)

	// Parse the terminal or expression on the right.
	right := parseSubOperation(p)

	// Recursively parse the expression on the right until an operator with
	// precedence less or equal to this one is encountered.
	right = parseOperation(p, right, op.Precedence())

	left = st.Operation{left, op, right}

	// Parse the remaining operations in this expression.
	left = parseOperation(p, left, leftPriority)

	return left
}

func parseSubOperation(p *pipe) st.Expression {
	// pattern := func_call | literal | group

	switch {
	case isFuncCall(p):
		return parseFuncCall(p)

	case isLiteral(p):
		return parseLiteral(p)

	case isGroup(p):
		return parseGroup(p)
	}

	panic(unexpected("parseRightSide", p.peek(), `function_call | literal | group`))
}

func isLiteral(p *pipe) bool {
	return p.matchAny(
		token.ID, // Yes I know, need to sort it out
		token.VOID,
		token.BOOL,
		token.NUMBER,
		token.STRING,
		token.TEMPLATE,
	)
}

func parseLiteral(p *pipe) st.Expression {
	tk := p.next()

	if tk.Type == token.ID {
		return st.Identifier(tk)
	} else {
		return st.Value(tk)
	}
}

func isGroup(p *pipe) bool {
	return p.match(token.PAREN_OPEN)
}

func parseGroup(p *pipe) st.Expression {
	// pattern := PAREN_OPEN expression PAREN_CLOSE

	p.expect(`group`, token.PAREN_OPEN)

	g := parseExpression(p)
	if g == nil {
		panic(unexpected("group", p.past(), token.ANOTHER))
	}

	p.expect(`group`, token.PAREN_CLOSE)
	return g
}
