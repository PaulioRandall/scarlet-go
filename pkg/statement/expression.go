package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Expression interface {
	Token() token.Token

	String(indent int) string
}

type Identifier struct {
	Fixed  bool
	Source token.Token
}

func (id Identifier) Token() token.Token {
	return id.Source
}

func (id Identifier) String(i int) string {

	var s str

	return s.indent(i).
		append("[Identifier] ").
		appendIf(id.Fixed, "(FIXED) ").
		append(id.Source.String()).
		String()
}

type Value struct {
	Source token.Token
}

func (v Value) Token() token.Token {
	return v.Source
}

func (v Value) String(i int) string {

	var s str

	return s.indent(i).
		append("[Value] ").
		append(v.Source.String()).
		String()
}

type List struct {
	Key   token.Token
	Open  token.Token
	Exprs []Expression
	Close token.Token
}

func (l List) Token() token.Token {
	return l.Key
}

func (l List) String(i int) string {

	var s str

	s.indent(i).
		append("[List] ").
		append(l.Key.String())

	s.newline().
		indent(i + 1).
		append("Open: ").
		append(l.Open.String())

	s.newline().
		indent(i + 1).
		append("Exprs:")

	s.newline().
		appendExps(i+2, l.Exprs)

	s.newline().
		indent(i + 1).
		append("Close: ").
		append(l.Close.String())

	return s.String()
}

type ListAccess struct {
	ID    Identifier
	Index Expression
}

func (la ListAccess) Token() token.Token {
	return la.ID.Source
}

func (la ListAccess) String(i int) string {

	var s str

	s.indent(i).
		append("[ListAccess] ").
		append(la.ID.Source.String())

	s.newline().
		indent(i + 1).
		append("Index:")

	s.newline().
		append(la.Index.String(i + 2))

	return s.String()
}
