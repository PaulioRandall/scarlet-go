package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Guard struct {
	Open  token.Token
	Cond  Expression
	Close token.Token
	Block Block
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
		append(g.Cond.String(i + 2))

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
