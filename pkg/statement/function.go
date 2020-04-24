package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type FuncDef struct {
	Open   token.Token
	Input  []token.Token
	Output []token.Token
	Body   Block
}

func (f FuncDef) Token() token.Token {
	return f.Open
}

func (f FuncDef) String(i int) string {

	var s str

	s.indent(i).
		append("[FuncDef] ").
		append(f.Open.String())

	s.newline().
		indent(i + 1).
		append("Input:")

	s.newline().
		appendTks(i+2, f.Input)

	s.newline().
		indent(i + 1).
		append("Output:")

	s.newline().
		appendTks(i+2, f.Output)

	s.newline().
		indent(i + 1).
		append("Body:")

	s.newline().
		append(f.Body.String(i + 2))

	return s.String()
}

type FuncCall struct {
	ID    Expression
	Input []Expression
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
		append("Input:")

	s.newline().
		appendExps(i+2, f.Input)

	return s.String()
}
