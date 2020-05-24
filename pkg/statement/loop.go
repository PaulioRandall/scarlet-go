package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Loop struct {
	Open     Token
	IndexVar Token
	Guard    Guard
}

func (l Loop) Token() Token {
	return l.Open
}

func (l Loop) String(i int) string {

	var s str

	s.indent(i).
		append("[Loop] ").
		appendTk(l.Open)

	s.newline().
		indent(i + 1).
		append("Index: ").
		appendTk(l.IndexVar)

	s.newline().
		indent(i + 1).
		append("Guard:")

	s.newline().
		append(l.Guard.String(i + 2))

	return s.String()
}
