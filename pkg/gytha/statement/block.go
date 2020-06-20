package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

type Block struct {
	Open  Token
	Stats []Statement
	Close Token
}

func (b Block) Token() Token {
	return b.Open
}

func (b Block) String(i int) string {

	var s str

	s.indent(i).
		append("[Block]")

	s.newline().
		indent(i + 1).
		append("Open: ").
		appendTk(b.Open)

	s.newline().
		indent(i + 1).
		append("Statements:")

	s.newline().
		appendStats(i+2, b.Stats)

	s.newline().
		indent(i + 1).
		append("Close: ").
		appendTk(b.Close)

	return s.String()
}
