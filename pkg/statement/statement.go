package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Print(stats []Statement) {

	var s str

	s.append("[Statements]").
		newline()
	s.appendStats(1, stats)
	s.print()
}

type Statement interface {
	Token() token.Token

	String(indent int) string
}

type Assignment struct {
	IDs    []Identifier
	Assign token.Token
	Exprs  []Expression
}

func (a Assignment) Token() token.Token {
	return a.Assign
}

func (a Assignment) String(indent int) string {

	var s str

	s.indent(indent).
		append("[Assignment] ").
		append(a.Assign.String())

	s.newline().
		indent(indent + 1).
		append("IDs:")

	s.newline().
		appendIds(indent+2, a.IDs)

	s.newline().
		indent(indent + 1).
		append("Exprs:")

	s.newline().
		appendExps(indent+2, a.Exprs)

	return s.String()
}
