package z_statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

type Expression interface {
	Token() Token
	String(indent int) string
}

type Value struct {
	tk Token
}

func (v Value) Token() Token {
	return Token(v.tk)
}

func (v Value) String(i int) string {

	var s str

	return s.indent(i).
		append("[Value] ").
		appendTk(v.tk).
		String()
}

type Identifier struct {
	tk Token
}

func (id Identifier) Token() Token {
	return Token(id.tk)
}

func (id Identifier) String(i int) string {

	var s str

	return s.indent(i).
		append("[Identifier] ").
		appendTk(id.tk).
		String()
}
