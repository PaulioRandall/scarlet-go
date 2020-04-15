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

type Arithmetic struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (a Arithmetic) Token() token.Token {
	return a.Operator
}

func (a Arithmetic) String(i int) string {

	str := indent(i) + "[Arithmetic] " + a.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += a.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += a.Right.String(i+2) + newline()

	return str
}

type Logic struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (l Logic) Token() token.Token {
	return l.Operator
}

func (l Logic) String(i int) string {

	str := indent(i) + "[Logic] " + l.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += l.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += l.Right.String(i+2) + newline()

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

// ExpressionOf returns either a Value or Identifier expression depending on the
// token type.
func ExpressionOf(tk token.Token) Expression {
	if tk.Type == token.ID {
		return Identifier{tk}
	}

	return Value{tk}
}
