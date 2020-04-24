package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Block struct {
	Open  token.Token
	Stats []Statement
	Close token.Token
}

func (b Block) Token() token.Token {
	return b.Open
}

func (b Block) String(i int) string {

	var s str

	s.indent(i).
		append("[Block]")

	s.newline().
		indent(i + 1).
		append("Open: ").
		append(b.Open.String())

	s.newline().
		indent(i + 1).
		append("Statements:")

	s.newline().
		appendStats(i+2, b.Stats)

	s.newline().
		indent(i + 1).
		append("Close: ").
		append(b.Close.String())

	return s.String()
}
