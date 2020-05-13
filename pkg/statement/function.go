package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type FuncDef struct {
	Key     token.Token
	Inputs  []token.Token
	Outputs []token.Token
	Body    Block
}

func (f FuncDef) Token() token.Token {
	return f.Key
}

func (f FuncDef) String(i int) string {

	var s str

	s.indent(i).
		append("[FuncDef] ").
		append(f.Key.String())

	s.newline().
		indent(i + 1).
		append("Inputs:")

	s.newline().
		appendTks(i+2, f.Inputs)

	s.newline().
		indent(i + 1).
		append("Outputs:")

	s.newline().
		appendTks(i+2, f.Outputs)

	s.newline().
		indent(i + 1).
		append("Body:")

	s.newline().
		append(f.Body.String(i + 2))

	return s.String()
}

type FuncCall struct {
	ID     Expression
	Inputs []Expression
}

func (f FuncCall) Token() token.Token {
	return f.ID.Token()
}

func (f FuncCall) String(i int) string {

	var s str

	s.indent(i).
		append("[FuncCall]")

	s.newline().
		indent(i + 1).
		append("ID:")

	s.newline().
		append(f.ID.String(i + 2))

	s.newline().
		indent(i + 1).
		append("Inputs:")

	s.newline().
		appendExps(i+2, f.Inputs)

	return s.String()
}
