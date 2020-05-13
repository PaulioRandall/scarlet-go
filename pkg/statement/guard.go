package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Guard struct {
	Open      token.Token
	Condition Expression
	Close     token.Token
	Block     Block
}

func (g Guard) Token() token.Token {
	return g.Open
}

func (g Guard) String(i int) string {

	var s str

	s.indent(i).
		append("[Guard]")

	s.newline().
		indent(i + 1).
		append("Open: ").
		append(g.Open.String())

	s.newline().
		indent(i + 1).
		append("Condition:")

	s.newline().
		append(g.Condition.String(i + 2))

	s.newline().
		indent(i + 1).
		append("Close: ").
		append(g.Close.String())

	s.newline().
		indent(i + 1).
		append("Statement:")

	s.newline().
		append(g.Block.String(i + 2))

	return s.String()
}
