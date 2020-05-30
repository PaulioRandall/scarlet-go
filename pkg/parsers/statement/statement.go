package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Statement interface {
	Expression
}

type Expression interface {
	Begin() (line, col int)
	End() (line, col int)
	String() string
}

type Identifier struct {
	TK Token
}

func (id Identifier) Begin() (int, int) {
	tk := id.TK
	return tk.Line(), tk.Col()
}

func (id Identifier) End() (int, int) {
	tk := id.TK
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (id Identifier) String() string {

	b := builder{}

	b.add(0, "[Identifier] ")
	b.addToken(0, id.TK)

	return b.String()
}

type Literal struct {
	TK Token
}

func (l Literal) Begin() (int, int) {
	return l.TK.Line(), l.TK.Col()
}

func (l Literal) End() (int, int) {
	tk := l.TK
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (l Literal) String() string {

	b := builder{}

	b.add(0, "[Literal] ")
	b.addToken(0, l.TK)

	return b.String()
}
