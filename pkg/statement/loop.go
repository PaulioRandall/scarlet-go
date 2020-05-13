package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Loop struct {
	Open     token.Token
	IndexVar token.Token
	Guard    Guard
}

func (l Loop) Token() token.Token {
	return l.Open
}

func (l Loop) String(i int) string {

	var s str

	s.indent(i).
		append("[Loop] ").
		append(l.Open.String())

	s.newline().
		indent(i + 1).
		append("Index: ").
		append(l.IndexVar.String())

	s.newline().
		indent(i + 1).
		append("Guard:")

	s.newline().
		append(l.Guard.String(i + 2))

	return s.String()
}
