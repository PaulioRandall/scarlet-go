package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

type AssignTarget struct {
	ID    Token
	Index Expression
}

func (at AssignTarget) Token() Token {
	return at.ID
}

func (at AssignTarget) String(i int) string {

	var s str

	s.newline().
		indent(i).
		append("[AssignTarget]")

	s.newline().
		indent(i + 1).
		append("ID: ").
		appendTk(at.ID)

	if at.Index != nil {
		s.newline().
			indent(i + 1).
			append("Index: ").
			newline().
			append(at.Index.String(i + 2))
	}

	return s.String()
}

type Assignment struct {
	Fixed   bool
	Targets []AssignTarget
	Assign  Token
	Exprs   []Expression
}

func (a Assignment) Token() Token {
	return a.Assign
}

func (a Assignment) String(i int) string {

	var s str

	s.indent(i).
		append("[Assignment] ").
		appendTk(a.Assign)

	s.newline().
		indent(i + 1).
		append("Targets:")

	appendAssignTargets(&s, i+2, a.Targets)

	s.newline().
		indent(i + 1).
		append("Exprs:")

	s.newline().
		appendExps(i+2, a.Exprs)

	return s.String()
}

func appendAssignTargets(s *str, indent int, ats []AssignTarget) *str {

	for _, at := range ats {
		s.append(at.String(indent))
	}

	return s
}
