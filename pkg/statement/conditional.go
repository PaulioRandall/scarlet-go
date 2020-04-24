package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Guard struct {
	Open  token.Token
	Cond  Expression
	Close token.Token
	Block Block
}

func (g Guard) Token() token.Token {
	return g.Open
}

func (g Guard) String(i int) string {

	s := indent(i) + "[Guard]" + newline()
	s += indent(i+1) + "Open: " + g.Open.String() + newline()
	s += indent(i+1) + "Condition:" + newline()
	s += g.Cond.String(i+2) + newline()
	s += indent(i+1) + "Close: " + g.Close.String() + newline()
	s += indent(i+1) + "Statement:" + newline()
	s += g.Block.String(i + 2)

	return s
}

type Match struct {
	Key   token.Token
	Open  token.Token
	Cases []Guard
	Close token.Token
}

func (m Match) Token() token.Token {
	return m.Key
}

func (m Match) String(i int) string {

	s := indent(i) + "[Match] " + m.Key.String() + newline()
	s += indent(i+1) + "Open: " + m.Open.String() + newline()

	s += indent(i+1) + "Cases:" + newline()
	for _, g := range m.Cases {
		s += g.String(i+2) + newline()
	}

	s += indent(i+1) + "Close: " + m.Close.String()

	return s
}
