package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// Expects the following token pattern:
// pattern := {operator operand}
// operand := literal | expression
func parseOperation(p *pipe, left st.Expression, leftPriority int) st.Expression {

	op := p.peek()

	if leftPriority >= st.Precedence(op.Type) {
		return left
	}

	p.expect(`parseOperation`, op.Type) // Because operator not taken yet

	right := parseRightSide(p)
	right = parseOperation(p, right, st.Precedence(op.Type))

	left = st.Operation{left, op, right}
	left = parseOperation(p, left, leftPriority)

	return left
}

// Expects the following token pattern:
// pattern := function | literal | group
func parseRightSide(p *pipe) st.Expression {

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
		token.ID,
		token.VOID,
		token.BOOL,
		token.NUMBER,
		token.STRING,
		token.TEMPLATE,
	)
}

// Expects the following token pattern:
// pattern := literal
func parseLiteral(p *pipe) st.Expression {
	tk := p.next()

	if tk.Type == token.ID {
		return st.Identifier{false, tk}
	} else {
		return st.Value{tk}
	}
}

func isGroup(p *pipe) bool {
	return p.match(token.PAREN_OPEN)
}

// Expects the following token pattern:
// pattern := PAREN_OPEN expression PAREN_CLOSE
func parseGroup(p *pipe) st.Expression {

	p.expect(`group`, token.PAREN_OPEN)

	g := parseExpression(p)
	if g == nil {
		panic(unexpected("group", p.past(), token.ANOTHER))
	}

	p.expect(`group`, token.PAREN_CLOSE)
	return g
}
