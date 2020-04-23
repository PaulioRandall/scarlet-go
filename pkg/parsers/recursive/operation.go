package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func parseOperation(p *pipe, left st.Expression, leftPriority int) st.Expression {

	op := p.snoop()

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

// Expects one of the following to appear next in the pipe:
// - Function
// - Literal value
// - Expression group
func parseRightSide(p *pipe) st.Expression {

	switch {
	case isFuncCall(p):
		return parseFuncCall(p)

	case isLiteral(p):
		return parseLiteral(p)

	case isGroup(p):
		return parseGroup(p)
	}

	panic(unexpected("parseRightSide", p.snoop(), `function_call | literal | group`))
}

func isLiteral(p *pipe) bool {
	return p.matchesNext(
		token.ID,
		token.VOID,
		token.BOOL,
		token.NUMBER,
		token.STRING,
		token.TEMPLATE,
	)
}

// Assumes isLiteral returns true.
func parseLiteral(p *pipe) st.Expression {
	tk := p.proceed()

	if tk.Type == token.ID {
		return st.Identifier{false, tk}
	} else {
		return st.Value{tk}
	}
}

func isGroup(p *pipe) bool {
	return p.inspect(token.PAREN_OPEN)
}

// Expects one of the following token patterns:
// - PAREN_OPEN, ..., PAREN_CLOSE
func parseGroup(p *pipe) st.Expression {

	p.expect(`group`, token.PAREN_OPEN)

	g := parseExpression(p)
	if g == nil {
		panic(unexpected("group", p.prior(), token.ANOTHER))
	}

	p.expect(`group`, token.PAREN_CLOSE)
	return g
}
