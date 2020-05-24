package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Match struct {
	Key   Token
	Open  Token
	Cases []Guard
	Close Token
}

func (m Match) Token() Token {
	return m.Key
}

func (m Match) String(i int) string {

	var s str

	s.indent(i).
		append("[Match] ").
		appendTk(m.Key)

	s.newline().
		indent(i + 1).
		append("Open: ").
		appendTk(m.Open)

	s.newline().
		indent(i + 1).
		append("Cases:")

	s.newline().
		appendGuards(i+2, m.Cases)

	s.newline().
		indent(i + 1).
		append("Close: ").
		appendTk(m.Close)

	return s.String()
}
