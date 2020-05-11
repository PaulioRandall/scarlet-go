package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type AssignTarget struct {
	ID    token.Token
	Index Expression
}

func (at AssignTarget) Token() token.Token {
	return at.ID
}

func (at AssignTarget) String(i int) string {

	var s str

	s.indent(i).
		append("[AssignTarget]")

	s.newline().
		indent(i).
		append("ID: ").
		append(at.ID.String())

	s.newline().
		indent(i).
		append("Index: ").
		newline().
		append(at.Index.String(i + 1))

	return s.String()
}
