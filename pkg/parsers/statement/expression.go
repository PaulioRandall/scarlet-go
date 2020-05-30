package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Expression interface {
	Begin() (line, col int)
	End() (line, col int)
	String() string
}

type Identifier struct {
	TK Token
}

func (id Identifier) Begin() (int, int) {
	return id.TK.Line(), id.TK.Col()
}

func (id Identifier) End() (int, int) {
	tk := id.TK
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (id Identifier) String() string {

	b := builder{}

	b.add("[Identifier] ")
	b.addToken(id.TK)

	return b.String()
}

type Value struct {
	TK Token
}

func (v Value) Begin() (int, int) {
	return v.TK.Line(), v.TK.Col()
}

func (v Value) End() (int, int) {
	tk := v.TK
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (v Value) String(i int) string {

	b := builder{}

	b.add("[Value] ")
	b.addToken(v.TK)

	return b.String()
}
