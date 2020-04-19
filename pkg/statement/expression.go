package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Expression interface {
	Token() token.Token

	String(indent int) string
}

type Identifier struct {
	Source token.Token
}

func (id Identifier) Token() token.Token {
	return id.Source
}

func (id Identifier) String(i int) string {
	return indent(i) + "[Identifier] " + id.Source.String()
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
		return Identifier{tk}
	default:
		return Value{tk}
	}
}

type List struct {
	Open  token.Token
	Exprs []Expression
	Close token.Token
}

func (l List) Token() token.Token {
	return l.Open
}

func (l List) String(i int) string {

	str := indent(i) + "[List]" + newline()
	str += indent(i+1) + "Open: " + l.Open.String() + newline()

	str += indent(i+1) + "Exprs: " + newline()
	for _, ex := range l.Exprs {
		str += ex.String(i+2) + newline()
	}

	str += indent(i+1) + "Close: " + l.Close.String()
	return str
}

type FuncCall struct {
	ID     token.Token
	Input  []token.Token
	Output []token.Token
}

func (f FuncCall) Token() token.Token {
	return f.ID
}

func (f FuncCall) String(i int) string {

	str := indent(i) + "[FuncCall] " + f.ID.String() + newline()

	str += indent(i+1) + "Input:" + newline()
	for _, tk := range f.Input {
		str += tk.String() + newline()
	}

	str += indent(i+1) + "Output: " + newline()
	for _, tk := range f.Input {
		str += tk.String() + newline()
	}

	return str
}
