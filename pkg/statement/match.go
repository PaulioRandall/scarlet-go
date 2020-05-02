package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Match struct {
	Key   token.Token
	Open  token.Token
	Cases []Guard
	Close token.Token
}

func (m Match) Token() token.Token {
	return m.Key
}

func (m Match) String(i int) string {

	var s str

	s.indent(i).
		append("[Match] ").
		append(m.Key.String())

	s.newline().
		indent(i + 1).
		append("Open: ").
		append(m.Open.String())

	s.newline().
		indent(i + 1).
		append("Cases:")

	s.newline().
		appendGuards(i+2, m.Cases)

	s.newline().
		indent(i + 1).
		append("Close: ").
		append(m.Close.String())

	return s.String()
}
