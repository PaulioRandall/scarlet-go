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

	str := indent(i) + "[Assignment]"
	if a.Assign != (token.Token{}) {
		str += " " + a.Assign.String()
	}
	str += newline()

	str += indent(i+1) + "IDs:" + newline()
	for _, tk := range a.IDs {
		str += indent(i+2) + tk.String() + newline()
	}

	str += indent(i+1) + "Exprs:"
	for _, expr := range a.Exprs {
		str += newline()
		str += expr.String(i + 2)
	}

	return str
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

	str := indent(i) + "[Guard] " + g.Open.String() + newline()
	str += g.Cond.String(i+1) + newline()
	str += indent(i+1) + g.Close.String() + newline()
	str += g.Stat.String(i + 1)

	return str
}

func indent(indent int) string {
	return strings.Repeat("\t", indent)
}

func newline() string {
	return "\n"
}
