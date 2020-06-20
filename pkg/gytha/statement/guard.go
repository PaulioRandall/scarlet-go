package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

type Guard struct {
	Open      Token
	Condition Expression
	Close     Token
	Block     Block
}

func (g Guard) Token() Token {
	return g.Open
}

func (g Guard) String(i int) string {

	var s str

	s.indent(i).
		append("[Guard]")

	s.newline().
		indent(i + 1).
		append("Open: ").
		appendTk(g.Open)

	s.newline().
		indent(i + 1).
		append("Condition:")

	s.newline().
		append(g.Condition.String(i + 2))

	s.newline().
		indent(i + 1).
		append("Close: ").
		appendTk(g.Close)

	s.newline().
		indent(i + 1).
		append("Statement:")

	s.newline().
		append(g.Block.String(i + 2))

	return s.String()
}
