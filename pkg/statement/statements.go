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

type Statement struct {
	IDs    []token.Token
	Assign token.Token
	Exprs  []Expression
}

func (s *Statement) String(i int) string {

	str := indent(i) + "[Statement]"
	if s.Assign != (token.Token{}) {
		str += " " + s.Assign.String()
	}
	str += newline()

	str += indent(i+1) + "IDs:" + newline()
	for _, tk := range s.IDs {
		str += indent(i+2) + tk.String() + newline()
	}

	str += indent(i+1) + "Exprs:"
	for _, expr := range s.Exprs {
		str += newline()
		str += expr.String(i + 2)
	}

	return str
}

func indent(indent int) string {
	return strings.Repeat("\t", indent)
}

func newline() string {
	return "\n"
}
