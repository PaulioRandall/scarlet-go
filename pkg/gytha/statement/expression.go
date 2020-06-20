package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

type Expression interface {
	Token() Token
	String(indent int) string
}

type Value struct {
	Tk Token
}

func (v Value) Token() Token {
	return Token(v.Tk)
}

func (v Value) String(i int) string {

	var s str

	return s.indent(i).
		append("[Value] ").
		appendTk(v.Tk).
		String()
}

type Identifier struct {
	Tk Token
}

func (id Identifier) Token() Token {
	return Token(id.Tk)
}

func (id Identifier) String(i int) string {

	var s str

	return s.indent(i).
		append("[Identifier] ").
		appendTk(id.Tk).
		String()
}

type Negation struct {
	Tk   Token
	Expr Expression
}

func (n Negation) Token() Token {
	return Token(n.Tk)
}

func (n Negation) String(i int) string {

	var s str

	s.indent(i).
		append("[Negation] ").
		appendTk(n.Tk)

	s.newline().
		indent(i + 1).
		append("Expr:")

	s.newline().
		append(n.Expr.String(i + 2))

	return s.String()
}
