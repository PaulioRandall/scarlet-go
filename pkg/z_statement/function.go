package z_statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

type FuncDef struct {
	Key     Token
	Inputs  []Token
	Outputs []Token
	Body    Block
}

func (f FuncDef) Token() Token {
	return f.Key
}

func (f FuncDef) String(i int) string {

	var s str

	s.indent(i).
		append("[FuncDef] ").
		appendTk(f.Key)

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

func (f FuncCall) Token() Token {
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
