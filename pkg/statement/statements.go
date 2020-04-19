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
	IDs    []token.Token
	Assign token.Token
	Exprs  []Expression
}

func (a Assignment) Token() token.Token {
	return a.Assign
}

func (a Assignment) String(i int) string {

	s := indent(i) + "[Assignment]"
	if a.Assign != (token.Token{}) {
		s += " " + a.Assign.String()
	}
	s += newline()

	s += indent(i+1) + "IDs:" + newline()
	for _, tk := range a.IDs {
		s += indent(i+2) + tk.String() + newline()
	}

	s += indent(i+1) + "Exprs:"
	for _, expr := range a.Exprs {
		s += newline()
		s += expr.String(i + 2)
	}

	return s
}

type Guard struct {
	Open  token.Token
	Cond  Expression
	Close token.Token
	Stat  Statement
}

func (g Guard) Token() token.Token {
	return g.Open
}

func (g Guard) String(i int) string {

	s := indent(i) + "[Guard]" + newline()
	s += indent(i+1) + "Open: " + g.Open.String() + newline()
	s += indent(i+1) + "Condition:" + newline()
	s += g.Cond.String(i+2) + newline()
	s += indent(i+1) + "Close: " + g.Close.String() + newline()
	s += indent(i+1) + "Statement:" + newline()
	s += g.Stat.String(i + 2)

	return s
}

func indent(indent int) string {
	return strings.Repeat("\t", indent)
}

func newline() string {
	return "\n"
}
