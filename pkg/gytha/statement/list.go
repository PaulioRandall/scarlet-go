package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

type List struct {
	Key   Token
	Open  Token
	Exprs []Expression
	Close Token
}

func (l List) Token() Token {
	return l.Key
}

func (l List) String(i int) string {

	var s str

	s.indent(i).
		append("[List] ").
		appendTk(l.Key)

	s.newline().
		indent(i + 1).
		append("Open: ").
		appendTk(l.Open)

	s.newline().
		indent(i + 1).
		append("Exprs:")

	s.newline().
		appendExps(i+2, l.Exprs)

	s.newline().
		indent(i + 1).
		append("Close: ").
		appendTk(l.Close)

	return s.String()
}

type ListAccess struct {
	ID    Identifier
	Index Expression
}

func (la ListAccess) Token() Token {
	return la.ID.Token()
}

func (la ListAccess) String(i int) string {

	var s str

	s.indent(i).
		append("[ListAccess] ").
		appendTk(la.ID.Token())

	s.newline().
		indent(i + 1).
		append("Index:")

	s.newline().
		append(la.Index.String(i + 2))

	return s.String()
}

type ListItemRef struct {
	Tk Token
}

func (r ListItemRef) Token() Token {
	return r.Tk
}

func (r ListItemRef) String(i int) string {

	var s str

	return s.indent(i).
		append("[ListItemRef] ").
		appendTk(r.Tk).
		String()
}
