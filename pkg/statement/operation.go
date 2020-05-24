package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Operation struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (op Operation) Token() Token {
	return op.Operator
}

func (op Operation) String(i int) string {

	var s str

	s.indent(i).
		append("[Operation] ").
		appendTk(op.Operator)

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
