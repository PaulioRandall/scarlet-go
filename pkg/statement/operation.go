package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Operation struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (op Operation) Token() token.Token {
	return op.Operator
}

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
