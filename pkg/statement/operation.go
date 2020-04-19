package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// Operation represents an mathematical operation, an expression with a left
// side, opertor, and right side.
type Operation struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

// Token satisfies the Expression interface.
func (op Operation) Token() token.Token {
	return op.Operator
}

// Precedence returns the priority of the expression type so it may be compared
// against other expression types. This is mostly useful for ordering oprations
// such as ensuring multiplications happen before additions.
func (op Operation) Precedence() int {
	return Precedence(op.Operator.Type)
}

// String satisfies the Expression interface.
func (op Operation) String(i int) string {

	str := indent(i) + "[Operation] " + op.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += op.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += op.Right.String(i + 2)

	return str
}

// IsBoolOperator returns true if tt represents an operator that produces a
// boolean value when evaluated.
func IsBoolOperator(tt token.TokenType) bool {
	switch tt {
	case token.LESS_THAN,
		token.LESS_THAN_OR_EQUAL,
		token.MORE_THAN,
		token.MORE_THAN_OR_EQUAL,
		token.EQUAL,
		token.NOT_EQUAL,
		token.AND,
		token.OR:

		return true
	}

	return false
}

// Precedence returns the precedences of the token type.
func Precedence(tt token.TokenType) int {
	switch tt {
	case token.MULTIPLY, token.DIVIDE, token.REMAINDER: // Multiplicative
		return 6
	case token.ADD, token.SUBTRACT: // Additive
		return 5
	case token.LESS_THAN, token.LESS_THAN_OR_EQUAL: // Relational
		fallthrough
	case token.MORE_THAN, token.MORE_THAN_OR_EQUAL: // Relational
		return 4
	case token.EQUAL, token.NOT_EQUAL: // Equality
		return 3
	case token.AND:
		return 2
	case token.OR:
		return 1
	default:
		return 0
	}
}
