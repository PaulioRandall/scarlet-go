package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type AssignTarget struct {
	ID    token.Token
	Index Expression
}

func (at AssignTarget) Token() token.Token {
	return at.ID
}

func (at AssignTarget) String(i int) string {

	var s str

	s.indent(i).
		append("[AssignTarget]")

	s.newline().
		indent(i).
		append("ID: ").
		append(at.ID.String())

	if at.Index != nil {
		s.newline().
			indent(i).
			append("Index: ").
			newline().
			append(at.Index.String(i + 1))
	}

	return s.String()
}

type Assignment struct {
	Fixed   bool
	Targets []AssignTarget
	Assign  token.Token
	Exprs   []Expression
}

func (a Assignment) Token() token.Token {
	return a.Assign
}

func (a Assignment) String(i int) string {

	var s str

	s.indent(i).
		append("[Assignment] ").
		append(a.Assign.String())

	s.newline().
		indent(i + 1).
		append("Targets:")

	s.newline().
		appendAssignTargets(i+2, a.Targets)

	s.newline().
		indent(i + 1).
		append("Exprs:")

	s.newline().
		appendExps(i+2, a.Exprs)

	return s.String()
}
