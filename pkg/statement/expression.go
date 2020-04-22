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

	s := indent(i) + "[Identifier] "

	if id.Fixed {
		s += "(FIXED) "
	}

	return s + id.Source.String()
}

type Value struct {
	Source token.Token
}

func (v Value) Token() token.Token {
	return v.Source
}

func (v Value) String(i int) string {
	return indent(i) + "[Value] " + v.Source.String()
}

// NewValueExpression returns either a Value or Identifier expression depending
// on the token type.
func NewValueExpression(tk token.Token) Expression {
	switch tk.Type {
	case token.ID:
		return Identifier{
			Source: tk,
		}
	default:
		return Value{tk}
	}
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

	str := indent(i) + "[List] " + l.Key.String() + newline()
	str += indent(i+1) + "Open: " + l.Open.String() + newline()

	str += indent(i+1) + "Exprs: " + newline()
	for _, ex := range l.Exprs {
		str += ex.String(i+2) + newline()
	}

	str += indent(i+1) + "Close: " + l.Close.String()
	return str
}

type ListAccess struct {
	ID    Identifier
	Index Expression
}

func (la ListAccess) Token() token.Token {
	return la.ID.Source
}

func (la ListAccess) String(i int) string {
	str := indent(i) + "[ListAccess] " + la.ID.Source.String() + newline()
	str += indent(i+1) + "Index: " + newline()
	return str + la.Index.String(i+2) + newline()
}
