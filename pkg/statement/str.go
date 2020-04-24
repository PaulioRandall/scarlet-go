package statement

import (
	"strings"
	//"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type str string

func (s *str) concat(texts ...string) {
	for _, txt := range texts {
		*s = str(*s + str(txt))
	}
}

func (s *str) indent(indent int) {
	s.concat(strings.Repeat("\t", indent))
}

func (s *str) newline() {
	s.concat("\n")
}

func (s *str) ids(ids []Identifier, indent int) {
	for i, id := range ids {
		if i != 0 {
			s.newline()
		}

		s.indent(indent)

		if id.Fixed {
			s.concat("FIXED ")
		}

		s.concat(id.Source.String())
	}
}

func indent(indent int) string {
	return strings.Repeat("\t", indent)
}

func newline() string {
	return "\n"
}
