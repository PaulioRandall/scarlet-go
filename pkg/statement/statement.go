package statement

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Print(ss []Statement) {

	s := "[Statements]"

	for _, st := range ss {
		s += newline()
		s += st.String(1)
	}

	s = strings.ReplaceAll(s, "\t", "  ")

	println(s)
	println(token.EOF)
	println()
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

func (a Assignment) String(i int) string {

	var s str

	s.indent(i)
	s.concat("[Assignment]")

	if a.Assign != (token.Token{}) {
		s.concat(" ", a.Assign.String())
	}

	s.newline()

	s.indent(i + 1)
	s.concat("IDs:")
	s.newline()

	s.ids(a.IDs, i+2)
	s.newline()

	s.indent(i + 1)
	s.concat("Exprs:")
	s.newline()
	s.exps(a.Exprs, i+2)

	return string(s)
}
