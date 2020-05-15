package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

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
	return la.ID.Token()
}

func (la ListAccess) String(i int) string {

	var s str

	s.indent(i).
		append("[ListAccess] ").
		append(la.ID.Token().String())

	s.newline().
		indent(i + 1).
		append("Index:")

	s.newline().
		append(la.Index.String(i + 2))

	return s.String()
}

type ListItemRef token.Token

func (r ListItemRef) Token() token.Token {
	return token.Token(r)
}

func (r ListItemRef) String(i int) string {

	var s str

	return s.indent(i).
		append("[ListItemRef] ").
		append(r.Token().String()).
		String()
}
