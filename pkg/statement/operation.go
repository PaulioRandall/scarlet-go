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

// String satisfies the Expression interface.
func (op Operation) String(i int) string {

	var s str

	s.indent(i).
		append("[Operation] ").
		append(op.Operator.String())

	s.newline().
		indent(i + 1).
		append("Left:")

	s.newline().
		append(op.Left.String(i + 2))

	s.newline().
		indent(i + 1).
		append("Right:")

	s.newline().
		append(op.Right.String(i + 2))

	return s.String()
}
