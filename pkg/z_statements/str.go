package z_statement

import (
	"fmt"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

type str string

func (s *str) String() string {
	return string(*s)
}

func (s *str) print() {
	fmt.Println(string(*s))
	fmt.Println()
}

func (s *str) append(txt string) *str {
	*s = str(*s + str(txt))
	return s
}

func (s *str) appendTk(tk Token) *str {
	*s = str(*s + str(ToString(tk)))
	return s
}

func (s *str) appendIf(condition bool, txt string) *str {
	if condition {
		return s.append(txt)
	}
	return s
}

func (s *str) indent(indent int) *str {
	return s.append(strings.Repeat("\t", indent))
}

func (s *str) newline() *str {
	return s.append("\n")
}

func (s *str) appendStats(indent int, stats []Statement) *str {

	for i, st := range stats {
		if i != 0 {
			s.newline()
		}

		s.append(st.String(indent))
	}

	return s
}

func (s *str) appendGuards(indent int, guards []Guard) *str {

	for i, g := range guards {
		if i != 0 {
			s.newline()
		}

		s.append(g.String(indent))
	}

	return s
}

func (s *str) appendExps(indent int, exps []Expression) *str {

	for i, exp := range exps {
		if i != 0 {
			s.newline()
		}

		s.append(exp.String(indent))
	}

	return s
}

func (s *str) appendIds(indent int, ids []Identifier) *str {

	for i, id := range ids {
		if i != 0 {
			s.newline()
		}

		s.appendTk(id.Token())
	}

	return s
}

func (s *str) appendTks(indent int, tks []Token) *str {

	for i, tk := range tks {
		if i != 0 {
			s.newline()
		}

		s.indent(indent).
			appendTk(tk)
	}

	return s
}

func indent(indent int) string {
	return strings.Repeat("\t", indent)
}

func newline() string {
	return "\n"
}
